package gochatwork

import (
	"fmt"
	"net/url"
)

// PostMessage post message to rooms/room_id/messages and response by int64
func (c *Client) PostMessage(roomID int64, message string) (int64, error) {
	var responseJSON = struct {
		MessageID int64 `json:"message_id"`
	}{}

	b, err := c.PostMessageRaw(roomID, message)
	err = setStructFromJSON(b, &responseJSON, err)
	return responseJSON.MessageID, err
}

// PostMessageRaw post message to rooms/room_id/messages and response by []byte
func (c *Client) PostMessageRaw(roomID int64, message string) ([]byte, error) {
	params := url.Values{}
	params.Add("body", message)
	return c.connection.Post(fmt.Sprintf("rooms/%d/messages", roomID), params, c.config)
}

// GetMessage get message to rooms/room_id/messages and response by []Message
func (c *Client) GetMessage(roomID int64, force bool) ([]Message, error) {
	var messages []Message

	b, err := c.GetMessageRaw(roomID, force)
	err = setStructFromJSON(b, &messages, err)
	return messages, err
}

// GetMessageRaw get message to rooms/room_id/messages and response by []byte
func (c *Client) GetMessageRaw(roomID int64, force bool) ([]byte, error) {
	params := url.Values{}
	if force {
		params.Add("force", "1")
	}

	return c.connection.Get(fmt.Sprintf("rooms/%d/messages", roomID), params, c.config)
}

// GetSpecificMessage get message to rooms/room_id/messages/message_id and response by Message
func (c *Client) GetSpecificMessage(roomID int64, messageID int64) (Message, error) {
	var message Message

	b, err := c.GetSpecificMessageRaw(roomID, messageID)
	err = setStructFromJSON(b, &message, err)
	return message, err
}

// GetSpecificMessageRaw get message to rooms/room_id/messages/message_id and response by []byte
func (c *Client) GetSpecificMessageRaw(roomID int64, messageID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d/messages/%d", roomID, messageID), url.Values{}, c.config)
}
