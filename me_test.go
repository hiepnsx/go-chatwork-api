package gochatwork

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMe(t *testing.T) {
	Convey("correct", t, func() {
		correctJSON := `{
		"account_id":42,
		"room_id":4242,
		"name":"name_text",
		"chatwork_id":"chatwork_id_text",
		"organization_id":424242,
		"organization_name":"organization_name_text",
		"department":"department_text",
		"title":"title_text",
		"url":"url_text",
		"introduction":"introduction_text",
		"mail":"mail_text",
		"tel_organization":"tel_organization_text",
		"tel_extension":"tel_extension_text",
		"tel_mobile":"tel_mobile_text",
		"skype":"skype_text",
		"facebook":"facebook_text",
		"twitter":"twitter_text",
		"avatar_image_url":"avatar_image_url_text"
		}`

		testToken := "testToken"
		client := New(testToken)

		Convey("MeRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = make([]byte, 0)
			client.connection = stub

			b, _ := client.MeRaw()
			So(len(b), ShouldEqual, 0)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "me")
		})

		Convey("Me", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			me, err := client.Me()
			So(err, ShouldBeNil)

			So(me.AccountID, ShouldEqual, 42)
			So(me.RoomID, ShouldEqual, 4242)
			So(me.Name, ShouldEqual, "name_text")
			So(me.ChatworkID, ShouldEqual, "chatwork_id_text")
			So(me.OrganizationID, ShouldEqual, 424242)
			So(me.OrganizationName, ShouldEqual, "organization_name_text")
			So(me.Department, ShouldEqual, "department_text")
			So(me.Title, ShouldEqual, "title_text")
			So(me.URL, ShouldEqual, "url_text")
			So(me.Introduction, ShouldEqual, "introduction_text")
			So(me.Mail, ShouldEqual, "mail_text")
			So(me.TelOrganization, ShouldEqual, "tel_organization_text")
			So(me.TelExtension, ShouldEqual, "tel_extension_text")
			So(me.TelMobile, ShouldEqual, "tel_mobile_text")
			So(me.Skype, ShouldEqual, "skype_text")
			So(me.Facebook, ShouldEqual, "facebook_text")
			So(me.Twitter, ShouldEqual, "twitter_text")
			So(me.AvatarImageURL, ShouldEqual, "avatar_image_url_text")
		})
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
