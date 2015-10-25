package gochatwork

import (
	"fmt"
	"net/url"
	"strings"
)

// GetMembers return rooms/room_id/members response by []Account
func (c *Client) GetMembers(roomID int64) ([]Account, error) {
	var accounts []Account

	b, err := c.GetMembersRaw(roomID)
	err = setSturctFromJSON(b, &accounts, err)
	return accounts, err
}

// GetMembersRaw return rooms/room_id/members response by []byte
func (c *Client) GetMembersRaw(roomID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d/members", roomID), url.Values{}, c.config)
}

// PutMembersRaw return PUT rooms/room_id/members response by []byte
func (c *Client) PutMembersRaw(roomID int64, membersAdminIDs []int64, membersMemberIDs []int64, membersReadonlyIDs []int64) ([]byte, error) {
	params := url.Values{}

	if len(membersAdminIDs) != 0 {
		str := fmt.Sprintf("%v", membersAdminIDs)
		str = strings.Trim(str, "[]")
		str = strings.Replace(str, " ", ",", -1)
		params.Add("members_admin_ids", str)
	}

	if len(membersMemberIDs) != 0 {
		str := fmt.Sprintf("%v", membersMemberIDs)
		str = strings.Trim(str, "[]")
		str = strings.Replace(str, " ", ",", -1)
		params.Add("members_member_ids", str)
	}

	if len(membersReadonlyIDs) != 0 {
		str := fmt.Sprintf("%v", membersReadonlyIDs)
		str = strings.Trim(str, "[]")
		str = strings.Replace(str, " ", ",", -1)
		params.Add("members_readonly_ids", str)
	}

	return c.connection.Put(fmt.Sprintf("rooms/%d/members", roomID), params, c.config)
}
