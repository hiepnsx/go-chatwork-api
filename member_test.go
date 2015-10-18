package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMembers(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		Convey("MembersRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte("")
			client.connection = stub

			b, _ := client.MembersRaw(42)
			So(string(b), ShouldEqual, "")
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/members")
		})
	})
}
