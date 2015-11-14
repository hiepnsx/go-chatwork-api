package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func CheckAccount(v *TestValue, account Account) {
	So(account.AccountID, ShouldEqual, v.GetInt64())
	So(account.Role, ShouldEqual, v.GetString())
	So(account.Name, ShouldEqual, v.GetString())
	So(account.ChatworkID, ShouldEqual, v.GetString())
	So(account.OrganizationID, ShouldEqual, v.GetInt64())
	So(account.OrganizationName, ShouldEqual, v.GetString())
	So(account.Department, ShouldEqual, v.GetString())
	So(account.AvatarImageURL, ShouldEqual, v.GetString())
}

func TestGetContacts(t *testing.T) {
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
	"avatar_image_url":"string_8"
  },{
	"account_id":9,
	"role":"string_10",
	"name":"string_11",
	"chatwork_id":"string_12",
	"organization_id":13,
	"organization_name":"string_14",
	"department":"string_15",
	"avatar_image_url":"string_16"
  }
]
`
		Convey("GetMembers", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			accounts, err := client.GetContacts()

			So(err, ShouldBeNil)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "contacts")

			So(len(accounts), ShouldEqual, 2)

			v := &TestValue{}
			v.Count = 1
			CheckAccount(v, accounts[0])
			CheckAccount(v, accounts[1])
		})

		Convey("GetContactsRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.GetContactsRaw()
			So(string(b), ShouldEqual, correctJSON)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "contacts")
		})
	})
}
