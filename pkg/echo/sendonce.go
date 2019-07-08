/*
Copyright © 2019 ARTEM KOLOMYTSEV kolomytsev1996@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
