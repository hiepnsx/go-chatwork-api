package main

import (
	"fmt"
	"os"

	chatwork "github.com/ota42y/go-chatwork-api"
)

func me() {
	token := os.Getenv("CHATWORK_API_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := chatwork.New(token)
	me, err := client.Me()
	if err == nil {
		fmt.Println(me)
	} else {
		fmt.Println(err)
	}
}
