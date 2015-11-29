[![Circle CI](https://circleci.com/gh/ota42y/go-chatwork-api.svg?style=svg)](https://circleci.com/gh/ota42y/go-chatwork-api)
[![GoDoc](https://godoc.org/github.com/ota42y/go-chatwork-api?status.svg)](https://godoc.org/github.com/ota42y/go-chatwork-api)

Chatwork api client for golang.

This support www.chatwork.com and kcw.kddi.ne.jp 

# Install
```bash
$ go get github.com/ota42y/go-chatwork-api
```

# Usage

```go
package main

import (
	chatwork "github.com/ota42y/go-chatwork-api"
)

func main() {
	chatwork := chatwork.New("api-key")
	
	rooms, err := client.GetRooms()
	if err == nil {
		for _, room := range rooms {
			fmt.Println(room.RoomId, room.Name, room.UnreadNum)
		}
	} else {
		fmt.Println(err)
	}
}
```


# Feature
Full api support. (in 2015/11)

- [x] /me
- [x] /my
- [x] /contacts
- [x] /rooms

