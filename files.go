package gochatwork

import (
	"fmt"
	"net/url"
)

// GetFiles return /rooms/{room_id}/files response by []File
func (c *Client) GetFiles(roomID int64, accountID int64) ([]File, error) {
	var files []File

	b, err := c.GetFilesRaw(roomID, accountID)
	err = setSturctFromJSON(b, &files, err)
	return files, err
}

// GetFilesRaw return /rooms/{room_id}/files response by []byte
func (c *Client) GetFilesRaw(roomID int64, accountID int64) ([]byte, error) {
	params := url.Values{}
	if accountID != 0 {
		params.Add("account_id", fmt.Sprintf("%d", accountID))
	}
	return c.connection.Get(fmt.Sprintf("rooms/%d/files", roomID), params, c.config)
}
