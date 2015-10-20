package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMembers(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `
[
  {
	"account_id":1,
	"role":"string_2",
	"name":"string_3",
	"chatwork_id":"string_4",
	"organization_id":5,
	"organization_name":"string_6",
	"department":"string_7",
	"avatar_image_url":"https://string_8"
  },{
	"account_id":9,
	"role":"string_10",
	"name":"string_11",
	"chatwork_id":"string_12",
	"organization_id":13,
	"organization_name":"string_14",
	"department":"string_15",
	"avatar_image_url":"https://string_16"
  }
]
`
		Convey("Members", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			accounts, err := client.Members(42)

			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/members")

			So(err, ShouldBeNil)
			So(len(accounts), ShouldEqual, 2)

			v := &TestValue{}
			v.Count = 1
			account := accounts[0]
			So(account.AccountID, ShouldEqual, v.GetInt64())
			So(account.Role, ShouldEqual, v.GetString())
			So(account.Name, ShouldEqual, v.GetString())
			So(account.ChatworkID, ShouldEqual, v.GetString())
			So(account.OrganizationID, ShouldEqual, v.GetInt64())
			So(account.OrganizationName, ShouldEqual, v.GetString())
			So(account.Department, ShouldEqual, v.GetString())
			So(account.AvatarImageURL, ShouldEqual, "https://"+v.GetString())

			account = accounts[1]
			So(account.AccountID, ShouldEqual, v.GetInt64())
			So(account.Role, ShouldEqual, v.GetString())
			So(account.Name, ShouldEqual, v.GetString())
			So(account.ChatworkID, ShouldEqual, v.GetString())
			So(account.OrganizationID, ShouldEqual, v.GetInt64())
			So(account.OrganizationName, ShouldEqual, v.GetString())
			So(account.Department, ShouldEqual, v.GetString())
			So(account.AvatarImageURL, ShouldEqual, "https://"+v.GetString())
		})

		Convey("MembersRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.MembersRaw(42)
			So(string(b), ShouldEqual, correctJSON)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/members")
		})
	})
}
