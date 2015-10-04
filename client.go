package gochatwork

import (
	"encoding/json"
	"net/url"
)

// Client is chatwork api client
type Client struct {
	config     *config
	connection apiConnection
}

type config struct {
	url   string
	token string
}

var chatworkURL = "https://api.chatwork.com"
var kddiChatworkURL = "https://kcw.kddi.ne.jp"

// New return api client for www.chatwork.com
func New(token string) *Client {
	return newClient(token, chatworkURL)
}

// NewKddiChatwork return api client for kcw.kddi.ne.jp
func NewKddiChatwork(token string) *Client {
	return newClient(token, kddiChatworkURL)
}

func newClient(token string, url string) *Client {
	c := &config{
		url:   url,
		token: token,
	}

	return &Client{
		config:     c,
		connection: &httpImp{},
	}
}

// Me is /me response struct
type Me struct {
	AccountID        int64 `json:"account_id"`
	RoomID           int64 `json:"room_id"`
	Name             string
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int64  `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string
	Title            string
	URL              string
	Introduction     string
	Mail             string
	TelOrganization  string `json:"tel_organization"`
	TelExtension     string `json:"tel_extension"`
	TelMobile        string `json:"tel_mobile"`
	Skype            string
	Facebook         string
	Twitter          string
	AvatarImageURL   string `json:"avatar_image_url"`
}

// Me return /me response by Me struct
func (c *Client) Me() (Me, error) {
	var me Me

	b, err := c.MeRaw()
	if err != nil {
		return me, err
	}

	err = json.Unmarshal(b, &me)
	return me, err
}

// MeRaw return /me response by []byte
func (c *Client) MeRaw() ([]byte, error) {
	return c.connection.Get("me", url.Values{}, c.config)
}

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
