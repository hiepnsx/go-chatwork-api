package gochatwork

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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

type stubHTTP struct {
	GetCount int
	GetByte  []byte
}

func (h *stubHTTP) Get() []byte {
	h.GetCount++
	return h.GetByte
}

func TestMe(t *testing.T) {
	testToken := "testToken"

	client := New(testToken)
	Convey("correct", t, func() {
		stub := &stubHTTP{}
		stub.GetByte = make([]byte, 0)
		client.http = stub

		So(len(client.MeRaw()), ShouldEqual, 0)
		So(stub.GetCount, ShouldEqual, 1)
	})
}
