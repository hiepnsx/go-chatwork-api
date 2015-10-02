package gochatwork

import (
	"encoding/json"
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

	b, err := c.connection.Get("me", c.config)
	if err != nil {
		return me, err
	}

	err = json.Unmarshal(b, &me)
	return me, err
}

// MeRaw return /me response by []byte
func (c *Client) MeRaw() ([]byte, error) {
	return c.connection.Get("me", c.config)
}

// MyStatusRaw return my/status response by []byte
func (c *Client) MyStatusRaw() ([]byte, error) {
	return c.connection.Get("my/status", c.config)
}