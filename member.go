package gochatwork

import (
	"fmt"
	"net/url"
)


// MembersRaw return rooms/room_id/members and response by []byte
func (c *Client) MembersRaw(roomID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d/members", roomID), url.Values{}, c.config)
}


