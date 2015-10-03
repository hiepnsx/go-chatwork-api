package main

import (
    "os"
    "fmt"

    chatwork "github.com/ota42y/go-chatwork-api"
)

func main() {
    token := os.Getenv("CHATWORK_API_TOKEN")
    if token == "" {
        fmt.Println("skip this test because no token")
        return
    }

    client := chatwork.New(token)
    status, err := client.MyStatus()
    if err == nil {
        fmt.Println(status)
    } else {
        fmt.Println(err)
    }
}
