package gochatwork

import (
	"encoding/json"
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
