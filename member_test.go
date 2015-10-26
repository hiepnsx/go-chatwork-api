package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetMembers(t *testing.T) {
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
		Convey("GetMembers", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			accounts, err := client.GetMembers(42)

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

		Convey("GetMembersRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.GetMembersRaw(42)
			So(string(b), ShouldEqual, correctJSON)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/members")
		})
	})
}

func TestPutMembers(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `{"admin":[1],"member":[2],"readonly":[3]}`
		Convey("PutMembers", func() {
			stub := &stubHTTP{}
			stub.PutByte = []byte(correctJSON)
			client.connection = stub

			membersAdminIDs := []int64{1}
			membersMemberIDs := []int64{2}
			membersReadonlyIDs := []int64{3}

			admin, member, readonly, err := client.PutMembers(42, membersAdminIDs, membersMemberIDs, membersReadonlyIDs)
			So(err, ShouldBeNil)
			So(stub.PutCount, ShouldEqual, 1)
			So(stub.PutEndPoint, ShouldEqual, "rooms/42/members")
			So(stub.PutParams.Get("members_admin_ids"), ShouldEqual, "1")
			So(stub.PutParams.Get("members_member_ids"), ShouldEqual, "2")
			So(stub.PutParams.Get("members_readonly_ids"), ShouldEqual, "3")

			So(admin, ShouldResemble, membersAdminIDs)
			So(member, ShouldResemble, membersMemberIDs)
			So(readonly, ShouldResemble, membersReadonlyIDs)
		})

		Convey("PutMembersRaw", func() {
			stub := &stubHTTP{}
			stub.PutByte = []byte(correctJSON)
			client.connection = stub

			membersAdminIDs := []int64{1, 2}
			membersMemberIDs := []int64{3}
			var membersReadonlyIDs []int64

			b, err := client.PutMembersRaw(42, membersAdminIDs, membersMemberIDs, membersReadonlyIDs)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, correctJSON)
			So(stub.PutCount, ShouldEqual, 1)
			So(stub.PutEndPoint, ShouldEqual, "rooms/42/members")
			So(stub.PutParams.Get("members_admin_ids"), ShouldEqual, "1,2")
			So(stub.PutParams.Get("members_member_ids"), ShouldEqual, "3")
			So(stub.PutParams.Get("members_readonly_ids"), ShouldEqual, "")
		})
	})
}
