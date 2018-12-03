package gockethook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Message to send
type Message struct {
	Text    string `json:"text"`
	Channel string `json:"channel,omitempty"`
}

// Connection to send messages
type Connection struct {
	url string
}

// New rocket hook
func New(url string) Connection {
	return Connection{url: url}
}

// SendText sends a text message.
func (conn Connection) SendText(msg string) error {
	return conn.Send(Message{Text: msg})
}

// Send sends a Message.
func (conn Connection) Send(msg Message) error {
	buffer, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	response, err := http.Post(conn.url, "application/json", bytes.NewReader(buffer))

	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf("send error: %d", response.StatusCode)
	}
	return nil
}
