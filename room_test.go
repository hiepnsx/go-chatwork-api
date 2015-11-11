package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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

			rooms, err := client.GetRooms()

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

			b, _ := client.GetRoomsRaw()
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

			room, err := client.GetSpecificRooms(41)
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

			b, _ := client.GetSpecificRoomsRaw(42)
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
		correctJSON := `{"room_id":42}`
		Convey("PutRooms", func() {
			stub := &stubHTTP{}
			stub.PutByte = []byte(correctJSON)
			client.connection = stub

			roomID, err := client.PutRooms(42, "desc", "meeting", "name")
			So(err, ShouldBeNil)
			So(stub.PutCount, ShouldEqual, 1)
			So(stub.PutEndPoint, ShouldEqual, "rooms/42")

			So(roomID, ShouldEqual, 42)
		})

		Convey("PutRoomsRaw", func() {
			stub := &stubHTTP{}
			stub.PutByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.PutRoomsRaw(42, "desc", "meeting", "name")
			So(string(b), ShouldEqual, correctJSON)
			So(stub.PutCount, ShouldEqual, 1)
			So(stub.PutEndPoint, ShouldEqual, "rooms/42")

			So(stub.PutParams.Get("description"), ShouldEqual, "desc")
			So(stub.PutParams.Get("icon_preset"), ShouldEqual, "meeting")
			So(stub.PutParams.Get("name"), ShouldEqual, "name")
		})
	})
}

func TestDeleteRooms(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := ``
		Convey("DeleteRooms", func() {
			stub := &stubHTTP{}
			stub.DeleteByte = []byte(correctJSON)
			client.connection = stub

			err := client.DeleteRooms(42, "leave")
			So(err, ShouldBeNil)
			So(stub.DeleteCount, ShouldEqual, 1)
			So(stub.DeleteEndPoint, ShouldEqual, "rooms/42")

			So(stub.DeleteParams.Get("action_type"), ShouldEqual, "leave")
		})
	})
}
