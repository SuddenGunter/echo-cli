package echo

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// SendOnce send message to websocket server, receives one response and then closes connection
// simple implementation, the only purpose of which is to fit initial task requirements
// client instance will be unusable after this method, but it's all done just for time-saving purposes
func SendOnce(client Client, token string, message string) (string, error) {
	err := client.Connect(token)
	if err != nil {
		return "", errors.Wrap(err, "Failed to connect to echo-server")
	}

	stdoutWriter := make(chan string)
	err = client.Listen(stdoutWriter)
	if err != nil {
		return "", errors.Wrap(err, "Failed to start listen on server")
	}

	err = client.SendMessage(message)
	if err != nil {
		return "", errors.Wrap(err, "Failed to send command to echo-server")
	}

	msg, ok := <-stdoutWriter
	var response string
	if ok {
		response = fmt.Sprintf("Server response: %v\n", msg)
	} else {
		log.Print("Connection closed")
	}

	err = client.Close()
	if err != nil {
		log.Fatal(err)
	}
	return response, nil
}
