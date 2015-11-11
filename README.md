
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
- [x] /me
- [x] /my
- [ ] /contacts
- [x] /rooms

