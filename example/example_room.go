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

	rooms, err := client.Rooms()
	if err == nil {
		fmt.Println(rooms)
	} else {
		fmt.Println(err)
	}
}
