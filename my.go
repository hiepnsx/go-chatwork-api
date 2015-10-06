package gochatwork

import (
	"encoding/json"
	"net/url"
)

// Status is /me response struct
type Status struct {
	UnreadRoomNum  int64 `json:"unread_room_num"`
	MentionRoomNum int64 `json:"mention_room_num"`
	MytaskRoomNum  int64 `json:"mytask_room_num"`
	UnreadNum      int64 `json:"unread_num"`
	MentionNum     int64 `json:"mention_num"`
	MytaskNum      int64 `json:"mytask_num"`
}

// MyStatus return my/status response by Status struct
func (c *Client) MyStatus() (Status, error) {
	var status Status

	b, err := c.MyStatusRaw()
	if err != nil {
		return status, err
	}

	err = json.Unmarshal(b, &status)
	return status, err
}

// MyStatusRaw return my/status response by []byte
func (c *Client) MyStatusRaw() ([]byte, error) {
	return c.connection.Get("my/status", url.Values{}, c.config)
}

// Room is room struct
type Room struct {
	RoomID   int64 `json:"room_id"`
	Name     string
	IconPath string `json:"icon_path"`
}

// Account is task struct
type Account struct {
	AccountID      int64 `json:"account_id"`
	Name           string
	AvatarImageURL string `json:"avatar_image_url"`
}

// Task is task struct
type Task struct {
	TaskID            int64 `json:"task_id"`
	Room              Room
	AssignedByAccount Account `json:"assigned_by_account"`
	MessageID         int64   `json:"message_id"`
	Body              string
	LimitTime         int64 `json:"limit_time"`
	Status            string
}

// MyTasks return my/tasks response by []Task
func (c *Client) MyTasks(params url.Values) ([]Task, error) {
	var tasks []Task

	b, err := c.MyTasksRaw(params)
	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(b, &tasks)
	return tasks, err
}

// MyTasksRaw return my/tasks response by []byte
func (c *Client) MyTasksRaw(params url.Values) ([]byte, error) {
	return c.connection.Get("my/tasks", params, c.config)
}
