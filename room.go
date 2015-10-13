package gochatwork

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Rooms return rooms response by []Room
func (c *Client) Rooms() ([]Room, error) {
	var rooms []Room

	b, err := c.RoomsRaw()
	if err != nil {
		return rooms, err
	}

	err = json.Unmarshal(b, &rooms)
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
	if err != nil {
		return room, err
	}

	err = json.Unmarshal(b, &room)
	return room, err
}

// RoomRaw return rooms/room_id response by []byte
func (c *Client) RoomRaw(roomID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d", roomID), url.Values{}, c.config)
}
