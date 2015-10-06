package gochatwork

import (
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMyStatus(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `{
		"unread_room_num":4,
		"mention_room_num":42,
		"mytask_room_num":424,
		"unread_num":4242,
		"mention_num":42424,
		"mytask_num":424242
		}`

		Convey("MyStatus", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			status, err := client.MyStatus()
			So(err, ShouldBeNil)
			So(status.UnreadRoomNum, ShouldEqual, 4)
			So(status.MentionRoomNum, ShouldEqual, 42)
			So(status.MytaskRoomNum, ShouldEqual, 424)
			So(status.UnreadNum, ShouldEqual, 4242)
			So(status.MentionNum, ShouldEqual, 42424)
			So(status.MytaskNum, ShouldEqual, 424242)
		})

		Convey("MyStatusRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = make([]byte, 0)
			client.connection = stub

			b, _ := client.MyStatusRaw()
			So(len(b), ShouldEqual, 0)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "my/status")
		})
	})
}

func TestMyTasks(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		Convey("MyTask", func() {
			correctJSON := `
[
  {
    "task_id":4,
    "room":{
      "room_id":42,
      "name":"room_name_1",
      "icon_path":"room_icon_path_1"
    },
    "assigned_by_account":{
      "account_id":424,
      "name":"assigned_by_account_name_1",
      "avatar_image_url":"assigned_by_account_avatar_image_url_1"
    },
    "message_id":4242,
    "body":"task_body_1",
    "limit_time":42424,
    "status":"done"
  },
  {
    "task_id":424242,
    "room":{
      "room_id":4242424,
      "name":"room_name_2",
      "icon_path":"room_icon_path_2"
    },
    "assigned_by_account":{
      "account_id":42424242,
      "name":"assigned_by_account_name_2",
      "avatar_image_url":"assigned_by_account_avatar_image_url_2"
    },
    "message_id":424242424,
    "body":"task_body_2",
    "limit_time":4242424242,
    "status":"open"
  }
]`
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			params := url.Values{}
			params.Add("assigned_by_account_id", "42")
			params.Add("status", "done")

			tasks, err := client.MyTasks(params)

			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "my/tasks")
			So(stub.GetParams, ShouldResemble, params)

			So(err, ShouldBeNil)
			So(len(tasks), ShouldEqual, 2)

			task := tasks[0]
			So(task.TaskID, ShouldEqual, 4)
			So(task.Room.RoomID, ShouldEqual, 42)
			So(task.Room.Name, ShouldEqual, "room_name_1")
			So(task.Room.IconPath, ShouldEqual, "room_icon_path_1")
			So(task.AssignedByAccount.AccountID, ShouldEqual, 424)
			So(task.AssignedByAccount.Name, ShouldEqual, "assigned_by_account_name_1")
			So(task.AssignedByAccount.AvatarImageURL, ShouldEqual, "assigned_by_account_avatar_image_url_1")
			So(task.MessageID, ShouldEqual, 4242)
			So(task.Body, ShouldEqual, "task_body_1")
			So(task.LimitTime, ShouldEqual, 42424)
			So(task.Status, ShouldEqual, "done")

			task = tasks[1]
			So(task.TaskID, ShouldEqual, 424242)
			So(task.Room.RoomID, ShouldEqual, 4242424)
			So(task.Room.Name, ShouldEqual, "room_name_2")
			So(task.Room.IconPath, ShouldEqual, "room_icon_path_2")
			So(task.AssignedByAccount.AccountID, ShouldEqual, 42424242)
			So(task.AssignedByAccount.Name, ShouldEqual, "assigned_by_account_name_2")
			So(task.AssignedByAccount.AvatarImageURL, ShouldEqual, "assigned_by_account_avatar_image_url_2")
			So(task.MessageID, ShouldEqual, 424242424)
			So(task.Body, ShouldEqual, "task_body_2")
			So(task.LimitTime, ShouldEqual, 4242424242)
			So(task.Status, ShouldEqual, "open")
		})

		Convey("MyTasksRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = make([]byte, 0)
			client.connection = stub

			params := url.Values{}
			params.Add("assigned_by_account_id", "42")
			params.Add("status", "done")

			b, _ := client.MyTasksRaw(params)
			So(len(b), ShouldEqual, 0)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "my/tasks")
			So(stub.GetParams, ShouldResemble, params)
		})
	})
}
