package main

import (
	"fmt"
	"os"

	chatwork "github.com/ota42y/go-chatwork-api"
)

func files() {
	token := os.Getenv("CHATWORK_API_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := chatwork.New(token)

	file, err := client.GetSpecificFiles(42, 21, true)
	if err == nil {
		fmt.Println(file)
	} else {
		fmt.Println(err)
	}

	files, err := client.GetFiles(42, 0)
	if err == nil {
		fmt.Println(files)
	} else {
		fmt.Println(err)
	}
}
