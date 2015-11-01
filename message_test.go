package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
