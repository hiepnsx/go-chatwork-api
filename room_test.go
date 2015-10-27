package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"strconv"
)

type TestValue struct {
	Count int64
}

func (t *TestValue) getCount() int64 {
	c := t.Count
	t.Count++
	return c
}

func (t *TestValue) GetInt64() int64 {
	return t.getCount()
}

func (t *TestValue) GetString() string {
	return "string_" + strconv.FormatInt(t.getCount(), 10)
}

func TestRooms(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		Convey("Rooms", func() {
			correctJSON := `
[
  {
    "room_id":1,
    "name":"string_2",
    "type":"my",
    "role":"member",
    "sticky":true,
    "unread_num":3,
    "mention_num":4,
    "mytask_num":5,
    "message_num":6,
    "file_num":7,
    "task_num":8,
    "icon_path":"string_9",
    "last_update_time":10
  },
  {
    "room_id":11,
    "name":"string_12",
    "type":"group",
    "role":"admin",
    "sticky":false,
    "unread_num":13,
    "mention_num":14,
    "mytask_num":15,
    "message_num":16,
    "file_num":17,
    "task_num":18,
    "icon_path":"string_19",
    "last_update_time":20
  }
]
`
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			rooms, err := client.Rooms()

			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms")

			So(err, ShouldBeNil)
			So(len(rooms), ShouldEqual, 2)

			v := &TestValue{}
			v.Count = 1
			room := rooms[0]
			So(room.RoomID, ShouldEqual, v.GetInt64())
			So(room.Name, ShouldEqual, v.GetString())
			So(room.Type, ShouldEqual, "my")
			So(room.Role, ShouldEqual, "member")
			So(room.Sticky, ShouldBeTrue)
			So(room.UnreadNum, ShouldEqual, v.GetInt64())
			So(room.MentionNum, ShouldEqual, v.GetInt64())
			So(room.MytaskNum, ShouldEqual, v.GetInt64())
			So(room.MessageNum, ShouldEqual, v.GetInt64())
			So(room.FileNum, ShouldEqual, v.GetInt64())
			So(room.TaskNum, ShouldEqual, v.GetInt64())
			So(room.IconPath, ShouldEqual, v.GetString())
			So(room.LastUpdateTime, ShouldEqual, v.GetInt64())

			room = rooms[1]
			So(room.RoomID, ShouldEqual, v.GetInt64())
			So(room.Name, ShouldEqual, v.GetString())
			So(room.Type, ShouldEqual, "group")
			So(room.Role, ShouldEqual, "admin")
			So(room.Sticky, ShouldBeFalse)
			So(room.UnreadNum, ShouldEqual, v.GetInt64())
			So(room.MentionNum, ShouldEqual, v.GetInt64())
			So(room.MytaskNum, ShouldEqual, v.GetInt64())
			So(room.MessageNum, ShouldEqual, v.GetInt64())
			So(room.FileNum, ShouldEqual, v.GetInt64())
			So(room.TaskNum, ShouldEqual, v.GetInt64())
			So(room.IconPath, ShouldEqual, v.GetString())
			So(room.LastUpdateTime, ShouldEqual, v.GetInt64())
		})

		Convey("RoomsRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = make([]byte, 0)
			client.connection = stub

			b, _ := client.RoomsRaw()
			So(len(b), ShouldEqual, 0)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms")
		})
	})
}

func TestRoom(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		Convey("Rooms", func() {
			correctJSON := `
{
  "room_id":1,
  "name":"string_2",
  "type":"direct",
  "role":"string_3",
  "sticky":true,
  "unread_num":4,
  "mention_num":5,
  "mytask_num":6,
  "message_num":7,
  "file_num":8,
  "task_num":9,
  "icon_path":"string_10",
  "description":"string_11",
  "last_update_time":12
}
`
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			room, err := client.Room(41)
			So(err, ShouldBeNil)

			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/41")

			v := &TestValue{}
			v.Count = 1
			So(room.RoomID, ShouldEqual, v.GetInt64())
			So(room.Name, ShouldEqual, v.GetString())
			So(room.Type, ShouldEqual, "direct")
			So(room.Role, ShouldEqual, v.GetString())
			So(room.Sticky, ShouldBeTrue)
			So(room.UnreadNum, ShouldEqual, v.GetInt64())
			So(room.MentionNum, ShouldEqual, v.GetInt64())
			So(room.MytaskNum, ShouldEqual, v.GetInt64())
			So(room.MessageNum, ShouldEqual, v.GetInt64())
			So(room.FileNum, ShouldEqual, v.GetInt64())
			So(room.TaskNum, ShouldEqual, v.GetInt64())
			So(room.IconPath, ShouldEqual, v.GetString())
			So(room.Description, ShouldEqual, v.GetString())
			So(room.LastUpdateTime, ShouldEqual, v.GetInt64())
		})

		Convey("RoomRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = make([]byte, 0)
			client.connection = stub

			b, _ := client.RoomRaw(42)
			So(len(b), ShouldEqual, 0)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42")
		})
	})
}

func TestPutRooms(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		Convey("PutRoomsRaw", func() {
			stub := &stubHTTP{}
			stub.PutByte = make([]byte, 0)
			client.connection = stub

			b, _ := client.PutRoomsRaw(42, "desc", "meeting", "name")
			So(len(b), ShouldEqual, 0)
			So(stub.PutCount, ShouldEqual, 1)
			So(stub.PutEndPoint, ShouldEqual, "rooms/42")

			So(stub.PutParams.Get("description"), ShouldEqual, "desc")
			So(stub.PutParams.Get("icon_preset"), ShouldEqual, "meeting")
			So(stub.PutParams.Get("name"), ShouldEqual, "name")
		})
	})
}

func TestPostMessage(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `{"message_id":42}`
		Convey("PostMessage", func() {
			stub := &stubHTTP{}
			stub.PostByte = []byte(correctJSON)
			client.connection = stub

			messageID, err := client.PostMessage(42, "test message")
			So(err, ShouldBeNil)
			So(stub.PostCount, ShouldEqual, 1)
			So(stub.PostEndPoint, ShouldEqual, "rooms/42/messages")
			So(stub.PostParams.Get("body"), ShouldEqual, "test message")

			So(messageID, ShouldEqual, 42)
		})

		Convey("PostMessageRaw", func() {
			stub := &stubHTTP{}
			stub.PostByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.PostMessageRaw(42, "test message")
			So(string(b), ShouldEqual, correctJSON)
			So(stub.PostCount, ShouldEqual, 1)
			So(stub.PostEndPoint, ShouldEqual, "rooms/42/messages")
			So(stub.PostParams.Get("body"), ShouldEqual, "test message")
		})
	})
}

func TestGetMessage(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `
[
  {
    "message_id":1,
    "account":{
      "account_id":2,
      "name":"string_3",
      "avatar_image_url":"https://string_4"
    },
    "body":"string_5",
    "send_time":6,
    "update_time":7
  }
]
`
		Convey("GetMessage", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			messages, err := client.GetMessage(42, true)

			So(err, ShouldBeNil)
			So(len(messages), ShouldEqual, 1)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/messages")
			So(stub.GetParams.Get("force"), ShouldEqual, "1")

			v := &TestValue{}
			v.Count = 1
			message := messages[0]
			So(message.MessageID, ShouldEqual, v.GetInt64())
			So(message.Account.AccountID, ShouldEqual, v.GetInt64())
			So(message.Account.Name, ShouldEqual, v.GetString())
			So(message.Account.AvatarImageURL, ShouldEqual, "https://"+v.GetString())
			So(message.Body, ShouldEqual, v.GetString())
			So(message.SendTime, ShouldEqual, v.GetInt64())
			So(message.UpdateTime, ShouldEqual, v.GetInt64())
		})

		Convey("GettMessageRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.GetMessageRaw(42, true)
			So(string(b), ShouldEqual, correctJSON)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/messages")
			So(stub.GetParams.Get("force"), ShouldEqual, "1")
		})
	})
}
