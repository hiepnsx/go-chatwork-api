package gochatwork

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
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
	GetCount    int
	GetByte     []byte
	GetEndPoint string
	GetParams   url.Values

	PostCount    int
	PostByte     []byte
	PostEndPoint string
	PostParams   url.Values
}

func (h *stubHTTP) Get(endPoint string, params url.Values, config *config) ([]byte, error) {
	h.GetCount++
	h.GetEndPoint = endPoint
	h.GetParams = params
	return h.GetByte, nil
}

func (h *stubHTTP) Post(endPoint string, params url.Values, config *config) ([]byte, error) {
	h.PostCount++
	h.PostEndPoint = endPoint
	h.PostParams = params
	return h.PostByte, nil
}
