package gochatwork

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Members return rooms/room_id/members response by []Account
func (c *Client) Members(roomID int64) ([]Account, error) {
	var accounts []Account

	b, err := c.MembersRaw(roomID)
	if err != nil {
		return accounts, err
	}

	err = json.Unmarshal(b, &accounts)
	return accounts, err
}

// MembersRaw return rooms/room_id/members response by []byte
func (c *Client) MembersRaw(roomID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d/members", roomID), url.Values{}, c.config)
}
