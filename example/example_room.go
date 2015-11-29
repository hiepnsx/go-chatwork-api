package main

import (
	"fmt"
	"os"

	chatwork "github.com/ota42y/go-chatwork-api"
)

func room() {
	token := os.Getenv("CHATWORK_API_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := chatwork.New(token)

	newRoomID, err := client.PostRooms("description", "project", []int64{}, []int64{}, []int64{}, "name")
	if err == nil {
		fmt.Println(newRoomID)
	} else {
		fmt.Println(err)
	}

	err = client.DeleteRooms(42, "delete")
	if err == nil {
		fmt.Println(err)
	}

	roomID, err := client.PutRooms(42, "dest", "idea", "test")
	if err == nil {
		fmt.Println(roomID)
	} else {
		fmt.Println(err)
	}

	rooms, err := client.GetRooms()
	if err == nil {
		fmt.Println(rooms)
	} else {
		fmt.Println(err)
	}

	room, err := client.GetSpecificRooms(42)
	if err == nil {
		fmt.Println(room)
	} else {
		fmt.Println(err)
	}
}
