package main

import (
	"fmt"
	"os"

	chatwork "github.com/ota42y/go-chatwork-api"
)

func member() {
	token := os.Getenv("CHATWORK_API_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := chatwork.New(token)
	admin := []int64{1}
	var member []int64
	read := []int64{2}
	ad, me, re, err := client.PutMembers(42, admin, member, read)
	if err == nil {
		fmt.Println(ad)
		fmt.Println(me)
		fmt.Println(re)
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
