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

	b, err := client.MembersRaw(42)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err)
	}
}
