package gochatwork

import (
	"net/url"
)

// GetContacts return GET /contacts response by []Account
func (c *Client) GetContacts() ([]Account, error) {
	var accounts []Account

	b, err := c.GetContactsRaw()
	err = setSturctFromJSON(b, &accounts, err)
	return accounts, err
}

// GetContactsRaw return GET /contacts response by []byte
func (c *Client) GetContactsRaw() ([]byte, error) {
	return c.connection.Get("contacts", url.Values{}, c.config)
}
