package gochatwork

import (
	"fmt"
	"net/url"
	"strings"
)

// GetRooms return rooms response by []Room
func (c *Client) GetRooms() ([]Room, error) {
	var rooms []Room

	b, err := c.GetRoomsRaw()
	err = setSturctFromJSON(b, &rooms, err)
	return rooms, err
}

// GetRoomsRaw return rooms response by []byte
func (c *Client) GetRoomsRaw() ([]byte, error) {
	return c.connection.Get("rooms", url.Values{}, c.config)
}

// PostRooms return POST rooms response by int64
func (c *Client) PostRooms(description string, iconPreset string, membersAdminIDs []int64, membersMemberIDs []int64, membersReadonlyIDs []int64, name string) (int64, error) {
	var responseJSON = struct {
		RoomID int64 `json:"room_id"`
	}{}

	b, err := c.PostRoomsRaw(description, iconPreset, membersAdminIDs, membersMemberIDs, membersReadonlyIDs, name)
	err = setSturctFromJSON(b, &responseJSON, err)
	return responseJSON.RoomID, err
}

// PostRoomsRaw return POST rooms response by []byte
func (c *Client) PostRoomsRaw(description string, iconPreset string, membersAdminIDs []int64, membersMemberIDs []int64, membersReadonlyIDs []int64, name string) ([]byte, error) {
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

	return c.connection.Post("rooms", params, c.config)
}

// GetSpecificRooms return rooms/room_id response by Room
func (c *Client) GetSpecificRooms(roomID int64) (Room, error) {
	var room Room

	b, err := c.GetSpecificRoomsRaw(roomID)
	err = setSturctFromJSON(b, &room, err)
	return room, err
}

// GetSpecificRoomsRaw return rooms/room_id response by []byte
func (c *Client) GetSpecificRoomsRaw(roomID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d", roomID), url.Values{}, c.config)
}

// PutRooms return PUT rooms/room_id response by int64
func (c *Client) PutRooms(roomID int64, description string, iconPreset string, name string) (int64, error) {
	var responseJSON = struct {
		RoomID int64 `json:"room_id"`
	}{}

	b, err := c.PutRoomsRaw(roomID, description, iconPreset, name)
	err = setSturctFromJSON(b, &responseJSON, err)
	return responseJSON.RoomID, err
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

// DeleteRooms send DELETE rooms/room_id response, this api don't return response
func (c *Client) DeleteRooms(roomID int64, actionType string) error {
	params := url.Values{}
	if actionType != "" {
		params.Add("action_type", actionType)
	}

	_, err := c.connection.Delete(fmt.Sprintf("rooms/%d", roomID), params, c.config)
	return err
}
