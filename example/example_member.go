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
	admin := []int64{1}
	member := []int64{2}
	var read []int64
	b, err := client.PutMembersRaw(42, admin, member, read)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println("error")
		fmt.Println(err)
	}

	accounts, err := client.GetMembers(42)
	if err == nil {
		fmt.Println(accounts)
	} else {
		fmt.Println("error")
		fmt.Println(err)
	}
}
