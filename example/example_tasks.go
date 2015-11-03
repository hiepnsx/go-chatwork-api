package main

import (
	"fmt"
	"os"

	chatwork "github.com/ota42y/go-chatwork-api"
)

func main() {
	token := os.Getenv("CHATWORK_API_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := chatwork.New(token)

	tasks, err := client.GetTasks(42, 0, 0, "done")
	if err == nil {
		fmt.Println(tasks)
	} else {
		fmt.Println(err)
	}

}
