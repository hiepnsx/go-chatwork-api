
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
	rooms, err := chatwork.Rooms()
	if err == nil {
		for _, room := range rooms {
			fmt.Println(room.RoomId, room.Name, room.UnreadNum)
		}
	}
}
```


# Feature
- [x] /me
- [x] /my
- [] /contacts
- [] /rooms
  - [x] GET /rooms
  - [] POST /rooms
  - [x] GET /rooms/{room_id}
  - [x] PUT /rooms/{room_id}
  - [x] DELETE /rooms/{room_id}
  - [x] GET /rooms/{room_id}/members
  - [x] PUT /rooms/{room_id}/members
  - [x] GET /rooms/{room_id}/messages
  - [x] POST /rooms/{room_id}/messages
  - [] GET /rooms/{room_id}/messages/{message_id}
  - [] GET /rooms/{room_id}/tasks
  - [] POST /rooms/{room_id}/tasks
  - [] GET /rooms/{room_id}/tasks/{task_id}
  - [] GET /rooms/{room_id}/files
  - [] GET /rooms/{room_id}/files{file_id}
  