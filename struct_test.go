package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetStructFromJSON(t *testing.T) {
	Convey("setSturctFromJSON", t, func() {
		json := `{"errors":["You don't have permission to get messages in this room"]}`

		var f File
		err := setStructFromJSON([]byte(json), &f, nil)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "[You don't have permission to get messages in this room]")
	})
}
