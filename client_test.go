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