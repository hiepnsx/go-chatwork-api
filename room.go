package gochatwork

import (
	"fmt"
	"net/url"
	"encoding/json"
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


// RoomRaw return rooms/room_id response by []byte
func (c *Client) RoomRaw(roomID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d", roomID), url.Values{}, c.config)
}
