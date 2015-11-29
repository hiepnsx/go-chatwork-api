package main

import (
	"fmt"
	"os"
	"time"

	chatwork "github.com/ota42y/go-chatwork-api"
)

func tasks() {
	token := os.Getenv("CHATWORK_API_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := chatwork.New(token)

	task, err := client.GetSpecificTask(42, 21)
	if err == nil {
		fmt.Println(task)
	} else {
		fmt.Println(err)
	}

	taskIDs, err := client.PostTasks(42, "new task", time.Now(), []int64{1})
	if err == nil {
		fmt.Println(taskIDs)
	} else {
		fmt.Println(err)
	}

	tasks, err := client.GetTasks(42, 0, 0, "done")
	if err == nil {
		fmt.Println(tasks)
	} else {
		fmt.Println(err)
	}
}
