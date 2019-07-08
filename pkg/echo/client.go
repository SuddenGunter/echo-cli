package echo

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/gorilla/websocket"
)

type Client interface {
	// Connect allows to establish websocket connection to configured server using user's auth token
	Connect(authToken string) error
	// SendMessage sends plain-text message over websocket channel to server
	SendMessage(message string) error
	// Listen start sending all received messages to configured messageReceiver
	Listen(chan string) error
	// Close sends websocket.CloseMessage to server and disconnects
	Close() error
}

type DefaultEchoServersClient struct {
	connectionString string
	route            string
	connection       *websocket.Conn
	closed           chan struct{}
}

// NewClient creates new client for websocket connection
// example config: connectionString: "localhost:8080", route: "/ws"
func NewClient(connectionString string, route string) *DefaultEchoServersClient {
	return &DefaultEchoServersClient{
		connectionString: connectionString,
		route:            route,
		closed:           make(chan struct{}),
	}
}

// Connect allows to establish websocket connection to configured server using user's auth token
func (client *DefaultEchoServersClient) Connect(authToken string) error {
	if client.connection != nil {
		return errors.WithStack(errors.New("Already connected"))
	}

	u := url.URL{Scheme: "ws", Host: client.connectionString, Path: client.route}
	header := http.Header{}
	header.Add("secretKey", authToken)
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return errors.Wrap(err, "Failed to establish websocket connection")
	}
	client.connection = conn
	return nil
}

// SendMessage sends plain-text message over websocket channel to server
func (client *DefaultEchoServersClient) SendMessage(message string) error {
	if client.connection == nil {
		return errors.WithStack(ErrNotConnected)
	}

	select {
	case <-client.closed:
		return errors.WithStack(errors.New("Connection closed"))
	default:
	}

	err := client.connection.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return errors.Wrap(err, "Failed to save message by websocket connection")
	}
	return nil
}

// Listen start sending all received messages to configured messageReceiver
func (client *DefaultEchoServersClient) Listen(messageReceiver chan string) error {
	if client.connection == nil {
		return errors.WithStack(ErrNotConnected)
	}

	go func() {
		for {
			select {
			case <-client.closed:
				return
			default:
			}

			_, message, err := client.connection.ReadMessage()
			if err != nil {
				log.Println("Failed to read message by websocket connection", err)
				return
			}
			messageReceiver <- string(message)
		}
	}()

	return nil
}

// Close sends websocket.CloseMessage to server and disconnects
func (client *DefaultEchoServersClient) Close() error {
	if client.connection == nil {
		return errors.WithStack(ErrNotConnected)
	}

	close(client.closed)
	//wait for readers/writers to finish receiving/sending messages
	time.Sleep(time.Second)

	err := client.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return errors.Wrap(err, "Failed to close websocket connection")
	}
	return nil
}

var ErrNotConnected = errors.New("Not connected to any server")
