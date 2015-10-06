package gochatwork

import (
    "testing"

    . "github.com/smartystreets/goconvey/convey"
)


func TestRooms(t *testing.T) {
    testToken := "testToken"
    client := New(testToken)

    Convey("correct", t, func() {

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
