package gochatwork

import (
	"net/url"
)

// Me return /me response by Me struct
func (c *Client) Me() (Me, error) {
	var me Me

	b, err := c.MeRaw()
	err = setStructFromJSON(b, &me, err)
	return me, err
}

// MeRaw return /me response by []byte
func (c *Client) MeRaw() ([]byte, error) {
	return c.connection.Get("me", url.Values{}, c.config)
}
