package main

import (
	"fmt"
	"os"

	chatwork "github.com/ota42y/go-chatwork-api"
)

func message() {
	token := os.Getenv("CHATWORK_API_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := chatwork.New(token)

	message, err := client.GetSpecificMessage(42, 21)
	if err == nil {
		fmt.Println(message)
	} else {
		fmt.Println(err)
	}

	messages, err := client.GetMessage(42, true)
	if err == nil {
		fmt.Println(messages)
	} else {
		fmt.Println(err)
	}

	messageID, err := client.PostMessage(42, "test")
	if err == nil {
		fmt.Println(messageID)
	} else {
		fmt.Println(err)
	}
}
