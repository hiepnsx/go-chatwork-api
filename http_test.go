package gochatwork

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestHttp(t *testing.T) {
	Convey("correct", t, func() {
		h := &httpImp{}

		_, err := h.Get("", url.Values{}, nil)
		So(err.Error(), ShouldEqual, "No auth token")
	})
}
