package gochatwork

import (
	"encoding/json"
	"net/url"
)

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
