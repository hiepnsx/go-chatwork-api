package gochatwork

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"os"
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

func (h *stubHTTP) Get(endPoint string, config *config) ([]byte, error) {
	h.GetCount++
	return h.GetByte, nil
}

func TestMe(t *testing.T) {
    Convey("correct", t, func() {
        testToken := "testToken"

        client := New(testToken)
        stub := &stubHTTP{}
		stub.GetByte = make([]byte, 0)
		client.connection = stub

        b, _ := client.MeRaw()

		So(len(b), ShouldEqual, 0)
		So(stub.GetCount, ShouldEqual, 1)
	})

    Convey("connect", t, func() {
        token := os.Getenv("CHATWORK_API_TOKEN")
		if token == "" {
			t.Log("skip this test because no token")
			return
		}

        client := New(token)
        b, err := client.MeRaw()
		So(len(b), ShouldNotEqual, 0)
		So(err, ShouldBeNil)
    })
}
