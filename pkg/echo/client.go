package echo

import (
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/gorilla/websocket"
)

type Client struct {
	connectionString string
	route            string
	authToken        string
	connection       *websocket.Conn
	closed           chan struct{}
	messageReceiver  chan string
}

func NewClient(connectionString string, route string, authToken string, messageReceiver chan string) *Client {
	return &Client{
		connectionString: connectionString,
		route:            route,
		authToken:        authToken,
		messageReceiver:  messageReceiver,
	}
}

func (client *Client) Connect() error {
	u := url.URL{Scheme: "ws", Host: client.connectionString, Path: client.route}
	header := http.Header{}
	header.Add("secretKey", client.authToken)
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return errors.Wrap(err, "Failed to establish websocket connection")
	}
	client.connection = conn
	return nil
}

func (client *Client) SendMessage(message string) error {
	//close???
	err := client.connection.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return errors.Wrap(err, "Failed to save message by websocket connection")
	}
	return nil
}

// Listen start sending all received messages to configured messageReceiver
func (client *Client) Listen() {
	go func() {
		defer close(client.closed)
		for {
			_, message, err := client.connection.ReadMessage()
			if err != nil {
				log.Println("Failed to read message by websocket connection", err)
				return
			}
			client.messageReceiver <- string(message)
		}
	}()
}

// Close sends websocket.CloseMessage to server and disconnects
func (client *Client) Close() error {
	err := client.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return errors.Wrap(err, "Failed to close websocket connection")
	}
	return nil
}
