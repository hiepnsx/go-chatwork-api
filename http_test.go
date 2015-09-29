package gochatwork

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHttp(t *testing.T) {
	Convey("correct", t, func() {
		h := &httpImp{}
		So(len(h.Get()), ShouldEqual, 0)
	})
}
