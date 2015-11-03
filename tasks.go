package gochatwork

import (
	"fmt"
	"net/url"
)

// GetTasks get tasks to rooms/room_id/tasks and response by []tasks
func (c *Client) GetTasks(roomID int64, accountID int64, assignedByAccountID int64, status string) ([]Task, error) {
	var tasks []Task

	b, err := c.GetTasksRaw(roomID, accountID, assignedByAccountID, status)
	err = setSturctFromJSON(b, &tasks, err)
	return tasks, err
}

// GetTasksRaw get tasks to rooms/room_id/tasks and response by []tasks
func (c *Client) GetTasksRaw(roomID int64, accountID int64, assignedByAccountID int64, status string) ([]byte, error) {
	params := url.Values{}
	if accountID != 0 {
		params.Add("account_id", fmt.Sprintf("%d", accountID))
	}

	if assignedByAccountID != 0 {
		params.Add("assigned_by_account_id", fmt.Sprintf("%d", assignedByAccountID))
	}

	if status != "" {
		params.Add("status", status)
	}

	return c.connection.Get(fmt.Sprintf("rooms/%d/tasks", roomID), params, c.config)
}
