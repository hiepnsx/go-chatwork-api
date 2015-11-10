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

// GetSpecificFiles return /rooms/{room_id}/files/file_id response by File
func (c *Client) GetSpecificFiles(roomID int64, fileID int64, createDownloadURL bool) (File, error) {
	var file File

	b, err := c.GetSpecificFilesRaw(roomID, fileID, createDownloadURL)
	err = setSturctFromJSON(b, &file, err)
	return file, err
}

// GetSpecificFilesRaw return /rooms/{room_id}/files/file_id response by []byte
func (c *Client) GetSpecificFilesRaw(roomID int64, fileID int64, createDownloadURL bool) ([]byte, error) {
	params := url.Values{}
	if createDownloadURL {
		params.Add("create_download_url", "1")
	}
	return c.connection.Get(fmt.Sprintf("rooms/%d/files/%d", roomID, fileID), params, c.config)
}
