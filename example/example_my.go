package main

import (
    "os"
    "fmt"

    chatwork "github.com/ota42y/go-chatwork-api"
    "net/url"
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

    params := url.Values{}
    params.Add("status", "done")

    b, err := client.MyTasksRaw(params)
    if err == nil {
        fmt.Println(string(b))
    } else {
        fmt.Println(err)
    }
}
