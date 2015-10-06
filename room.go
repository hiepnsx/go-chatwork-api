package gochatwork

import (
    "net/url"
)

// RoomsRaw return rooms response by []byte
func (c *Client) RoomsRaw() ([]byte, error) {
    return c.connection.Get("rooms", url.Values{}, c.config)
}
