package gochatwork

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHttp(t *testing.T) {
	Convey("correct", t, func() {
		h := &httpImp{}

        _, err := h.Get("", nil)
		So(err.Error(), ShouldEqual, "No auth token")
	})
}
