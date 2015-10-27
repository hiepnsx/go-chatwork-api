package gochatwork

import (
	"fmt"
	"net/url"
)

// Rooms return rooms response by []Room
func (c *Client) Rooms() ([]Room, error) {
	var rooms []Room

	b, err := c.RoomsRaw()
	err = setSturctFromJSON(b, &rooms, err)
	return rooms, err
}

// RoomsRaw return rooms response by []byte
func (c *Client) RoomsRaw() ([]byte, error) {
	return c.connection.Get("rooms", url.Values{}, c.config)
}

// Room return rooms/room_id response by Room
func (c *Client) Room(roomID int64) (Room, error) {
	var room Room

	b, err := c.RoomRaw(roomID)
	err = setSturctFromJSON(b, &room, err)
	return room, err
}

// RoomRaw return rooms/room_id response by []byte
func (c *Client) RoomRaw(roomID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d", roomID), url.Values{}, c.config)
}

// PutRoomsRaw return PUT rooms/room_id response by []byte
func (c *Client) PutRoomsRaw(roomID int64, description string, iconPreset string, name string) ([]byte, error) {
	params := url.Values{}
	if description != "" {
		params.Add("description", description)
	}

	if iconPreset != "" {
		params.Add("icon_preset", iconPreset)
	}

	if name != "" {
		params.Add("name", name)
	}

	return c.connection.Put(fmt.Sprintf("rooms/%d", roomID), params, c.config)
}

// PostMessage post message to rooms/room_id/messages and response by int64
func (c *Client) PostMessage(roomID int64, message string) (int64, error) {
	var responseJSON = struct {
		MessageID int64 `json:"message_id"`
	}{}

	b, err := c.PostMessageRaw(roomID, message)
	err = setSturctFromJSON(b, &responseJSON, err)
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
	err = setSturctFromJSON(b, &messages, err)
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
