package gochatwork

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
    testToken := "testToken"

    Convey("normal", t, func() {
        client := New(testToken)
        So(client.config.url, ShouldEqual, chatworkURL)
        So(client.config.token, ShouldEqual, testToken)
    })

    Convey("kddi", t, func() {
        client := NewKddiChatwork(testToken)
        So(client.config.url, ShouldEqual, kddiChatworkURL)
        So(client.config.token, ShouldEqual, testToken)
    })
}


type stubHttp struct {
    GetCount int
    GetByte []byte
}

func (h *stubHttp) Get() []byte {
    h.GetCount += 1
    return h.GetByte
}

func TestMe(t *testing.T) {
    testToken := "testToken"

    client := New(testToken)
    Convey("correct", t, func() {
        stub := &stubHttp{}
        stub.GetByte = make([]byte, 0)
        client.http = stub

        So(len(client.MeRaw()), ShouldEqual, 0)
        So(stub.GetCount, ShouldEqual, 1)
    })
}