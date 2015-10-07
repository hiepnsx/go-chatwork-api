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

			var r Room
			r.Type = "aaa"

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
