package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func CheckRoom(v *TestValue, room Room) {
	So(room.RoomID, ShouldEqual, v.GetInt64())
	So(room.Name, ShouldEqual, v.GetString())
	So(room.Type, ShouldEqual, v.GetString())
	So(room.Role, ShouldEqual, v.GetString())
	So(room.UnreadNum, ShouldEqual, v.GetInt64())
	So(room.MentionNum, ShouldEqual, v.GetInt64())
	So(room.MytaskNum, ShouldEqual, v.GetInt64())
	So(room.MessageNum, ShouldEqual, v.GetInt64())
	So(room.FileNum, ShouldEqual, v.GetInt64())
	So(room.TaskNum, ShouldEqual, v.GetInt64())
	So(room.IconPath, ShouldEqual, v.GetString())
	So(room.LastUpdateTime, ShouldEqual, v.GetInt64())
}

func TestGetRooms(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		Convey("GetRooms", func() {
			correctJSON := `
[
  {
    "room_id":1,
    "name":"string_2",
    "type":"string_3",
    "role":"string_4",
    "sticky":true,
    "unread_num":5,
    "mention_num":6,
    "mytask_num":7,
    "message_num":8,
    "file_num":9,
    "task_num":10,
    "icon_path":"string_11",
    "last_update_time":12
  },
  {
    "room_id":13,
    "name":"string_14",
    "type":"string_15",
    "role":"string_16",
    "sticky":false,
    "unread_num":17,
    "mention_num":18,
    "mytask_num":19,
    "message_num":20,
    "file_num":21,
    "task_num":22,
    "icon_path":"string_23",
    "last_update_time":24
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
			CheckRoom(v, room)
			So(room.Sticky, ShouldBeTrue)

			room = rooms[1]
			CheckRoom(v, room)
			So(room.Sticky, ShouldBeFalse)
		})

		Convey("GetRoomsRaw", func() {
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

func TestGetSpecificRooms(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		Convey("GetSpecificRooms", func() {
			correctJSON := `
{
  "room_id":1,
  "name":"string_2",
  "type":"string_3",
  "role":"string_4",
  "sticky":true,
  "unread_num":5,
  "mention_num":6,
  "mytask_num":7,
  "message_num":8,
  "file_num":9,
  "task_num":10,
  "icon_path":"string_11",
  "description":"description",
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
			CheckRoom(v, room)
			So(room.Description, ShouldEqual, "description")
			So(room.Sticky, ShouldBeTrue)
		})

		Convey("GetSpecificRoomsRaw", func() {
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

func TestPostRoomsRaw(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `{"room_id":42}`
		Convey("PostMessage", func() {
			stub := &stubHTTP{}
			stub.PostByte = []byte(correctJSON)
			client.connection = stub

			roomID, err := client.PostRooms("description", "project", []int64{1, 2}, []int64{3, 4}, []int64{5, 6}, "name")
			So(err, ShouldBeNil)
			So(stub.PostCount, ShouldEqual, 1)
			So(stub.PostEndPoint, ShouldEqual, "rooms")
			So(stub.PostParams.Get("description"), ShouldEqual, "description")
			So(stub.PostParams.Get("icon_preset"), ShouldEqual, "project")
			So(stub.PostParams.Get("members_admin_ids"), ShouldEqual, "1,2")
			So(stub.PostParams.Get("members_member_ids"), ShouldEqual, "3,4")
			So(stub.PostParams.Get("members_readonly_ids"), ShouldEqual, "5,6")
			So(stub.PostParams.Get("name"), ShouldEqual, "name")

			So(roomID, ShouldEqual, 42)
		})

		Convey("PostRoomsRaw", func() {
			stub := &stubHTTP{}
			stub.PostByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.PostRoomsRaw("description", "project", []int64{1, 2}, []int64{3, 4}, []int64{5, 6}, "name")
			So(string(b), ShouldEqual, correctJSON)
			So(stub.PostCount, ShouldEqual, 1)
			So(stub.PostEndPoint, ShouldEqual, "rooms")
			So(stub.PostParams.Get("description"), ShouldEqual, "description")
			So(stub.PostParams.Get("icon_preset"), ShouldEqual, "project")
			So(stub.PostParams.Get("members_admin_ids"), ShouldEqual, "1,2")
			So(stub.PostParams.Get("members_member_ids"), ShouldEqual, "3,4")
			So(stub.PostParams.Get("members_readonly_ids"), ShouldEqual, "5,6")
			So(stub.PostParams.Get("name"), ShouldEqual, "name")
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
