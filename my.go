package gochatwork

import (
	"net/url"
)

// MyStatus return my/status response by Status struct
func (c *Client) MyStatus() (Status, error) {
	var status Status

	b, err := c.MyStatusRaw()
	err = setStructFromJSON(b, &status, err)
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
	err = setStructFromJSON(b, &tasks, err)
	return tasks, err
}

// MyTasksRaw return my/tasks response by []byte
func (c *Client) MyTasksRaw(params url.Values) ([]byte, error) {
	return c.connection.Get("my/tasks", params, c.config)
}
