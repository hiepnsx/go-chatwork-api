package gochatwork

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestHttp(t *testing.T) {
    Convey("correct", t, func() {
        h := &httpImp{}
        So(len(h.Get()), ShouldEqual, 0)
    })
}
