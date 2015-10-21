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

	accounts, err := client.Members(42)
	if err == nil {
		fmt.Println(accounts)
	} else {
		fmt.Println("error")
		fmt.Println(err)
	}
}