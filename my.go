package gochatwork

import (
	"encoding/json"
	"net/url"
)

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
