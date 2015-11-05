package gochatwork

import (
	"fmt"
	"net/url"
	"strings"
	"time"
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

// PostTasks post tasks to rooms/room_id/tasks and response by []int64
func (c *Client) PostTasks(roomID int64, body string, limit time.Time, toIDs []int64) ([]int64, error) {
	var responseJSON = struct {
		TaskIDs []int64 `json:"task_ids"`
	}{}

	b, err := c.PostTasksRaw(roomID, body, limit, toIDs)
	err = setSturctFromJSON(b, &responseJSON, err)
	return responseJSON.TaskIDs, err
}

// PostTasksRaw post tasks to rooms/room_id/tasks and response by []byte
func (c *Client) PostTasksRaw(roomID int64, body string, limit time.Time, toIDs []int64) ([]byte, error) {
	params := url.Values{}
	if body != "" {
		params.Add("body", body)
	}

	if !limit.IsZero() {
		params.Add("limit", fmt.Sprintf("%d", limit.Unix()))
	}

	if len(toIDs) != 0 {
		str := fmt.Sprintf("%v", toIDs)
		str = strings.Trim(str, "[]")
		str = strings.Replace(str, " ", ",", -1)
		params.Add("to_ids", str)
	}

	return c.connection.Post(fmt.Sprintf("rooms/%d/tasks", roomID), params, c.config)
}

// GetSpecificTask get tasks to rooms/room_id/tasks/task_id and response by Task
func (c *Client) GetSpecificTask(roomID int64, taskID int64) (Task, error) {
	var task Task

	b, err := c.GetSpecificTaskRaw(roomID, taskID)
	err = setSturctFromJSON(b, &task, err)
	return task, err
}

// GetSpecificTaskRaw get tasks to rooms/room_id/tasks/task_id and response by []byte
func (c *Client) GetSpecificTaskRaw(roomID int64, taskID int64) ([]byte, error) {
	return c.connection.Get(fmt.Sprintf("rooms/%d/tasks/%d", roomID, taskID), url.Values{}, c.config)
}
