package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"time"
)

func CheckAccountLite(v *TestValue, a Account) {
	So(a.AccountID, ShouldEqual, v.GetInt64())
	So(a.Name, ShouldEqual, v.GetString())
	So(a.AvatarImageURL, ShouldEqual, v.GetString())
}

func CheckTask(v *TestValue, t Task) {
	So(t.TaskID, ShouldEqual, v.GetInt64())
	CheckAccountLite(v, t.Account)
	CheckAccountLite(v, t.AssignedByAccount)
	So(t.MessageID, ShouldEqual, v.GetInt64())
	So(t.Body, ShouldEqual, v.GetString())
	So(t.LimitTime, ShouldEqual, v.GetInt64())
	So(t.Status, ShouldEqual, "open")
}

func TestGetTasks(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `
[
  {
    "task_id": 1,
    "account": {
      "account_id": 2,
      "name": "string_3",
      "avatar_image_url": "string_4"
    },
    "assigned_by_account": {
      "account_id": 5,
      "name": "string_6",
      "avatar_image_url": "string_7"
    },
    "message_id": 8,
    "body": "string_9",
    "limit_time": 10,
    "status": "open"
  }
]
`
		Convey("GetTasks", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			tasks, err := client.GetTasks(42, 1, 2, "done")

			So(err, ShouldBeNil)
			So(len(tasks), ShouldEqual, 1)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/tasks")

			v := &TestValue{}
			v.Count = 1
			CheckTask(v, tasks[0])
		})

		Convey("GetTasksRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.GetTasksRaw(42, 1, 2, "done")
			So(string(b), ShouldEqual, correctJSON)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/tasks")
			So(stub.GetParams.Get("account_id"), ShouldEqual, "1")
			So(stub.GetParams.Get("assigned_by_account_id"), ShouldEqual, "2")
			So(stub.GetParams.Get("status"), ShouldEqual, "done")
		})
	})
}

func TestPostTasks(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `{"task_ids":[1, 2]}`
		Convey("GetTasks", func() {
			stub := &stubHTTP{}
			stub.PostByte = []byte(correctJSON)
			client.connection = stub

			now := time.Unix(100, 0)
			taskIds, err := client.PostTasks(42, "new task", now, []int64{1, 2, 3})
			So(err, ShouldBeNil)
			So(stub.PostCount, ShouldEqual, 1)
			So(stub.PostEndPoint, ShouldEqual, "rooms/42/tasks")
			So(stub.PostParams.Get("body"), ShouldEqual, "new task")
			So(stub.PostParams.Get("limit"), ShouldEqual, "100")
			So(stub.PostParams.Get("to_ids"), ShouldEqual, "1,2,3")

			So(taskIds, ShouldResemble, []int64{1, 2})
		})

		Convey("GetTasksRaw", func() {
			stub := &stubHTTP{}
			stub.PostByte = []byte(correctJSON)
			client.connection = stub

			now := time.Unix(100, 0)
			b, _ := client.PostTasksRaw(42, "new task", now, []int64{1, 2, 3})
			So(string(b), ShouldEqual, correctJSON)
			So(stub.PostCount, ShouldEqual, 1)
			So(stub.PostEndPoint, ShouldEqual, "rooms/42/tasks")
			So(stub.PostParams.Get("body"), ShouldEqual, "new task")
			So(stub.PostParams.Get("limit"), ShouldEqual, "100")
			So(stub.PostParams.Get("to_ids"), ShouldEqual, "1,2,3")
		})
	})
}

func TestGetSpecificTask(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `
{
  "status": "open",
  "limit_time": 10,
  "body": "string_9",
  "message_id": 8,
  "assigned_by_account": {
    "avatar_image_url": "string_7",
    "name": "string_6",
    "account_id": 5
  },
  "account": {
    "avatar_image_url": "string_4",
    "name": "string_3",
    "account_id": 2
  },
  "task_id": 1
}`
		Convey("GetTasks", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			task, err := client.GetSpecificTask(42, 21)
			So(err, ShouldBeNil)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/tasks/21")

			v := &TestValue{}
			v.Count = 1
			CheckTask(v, task)
		})

		Convey("GetSpecificTaskRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.GetSpecificTaskRaw(42, 21)
			So(string(b), ShouldEqual, correctJSON)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/tasks/21")
		})
	})
}
