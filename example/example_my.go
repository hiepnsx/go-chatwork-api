package main

import (
	"fmt"
	"os"

	chatwork "github.com/ota42y/go-chatwork-api"
	"net/url"
)

func my() {
	token := os.Getenv("CHATWORK_API_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := chatwork.New(token)

	status, err := client.MyStatus()
	if err == nil {
		fmt.Println(status)
	} else {
		fmt.Println(err)
	}

	params := url.Values{}
	params.Add("status", "done")

	tasks, err := client.MyTasks(params)
	if err == nil {
		fmt.Println(tasks)
	} else {
		fmt.Println(err)
	}
}
