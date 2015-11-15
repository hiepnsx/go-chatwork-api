package gochatwork

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func CheckFile(v *TestValue, f File) {
	So(f.FileID, ShouldEqual, v.GetInt64())
	CheckAccountLite(v, f.Account)
	So(f.MessageID, ShouldEqual, v.GetInt64())
	So(f.Filename, ShouldEqual, v.GetString())
	So(f.Filesize, ShouldEqual, v.GetInt64())
	So(f.UploadTime, ShouldEqual, v.GetInt64())
}

func TestGetFiles(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `
[
  {
    "file_id": 1,
    "account": {
      "account_id": 2,
      "name": "string_3",
      "avatar_image_url": "string_4"
    },
    "message_id": 5,
    "filename": "string_6",
    "filesize": 7,
    "upload_time": 8
  }
]
`
		Convey("GetFiles", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			files, err := client.GetFiles(42, 21)
			So(err, ShouldBeNil)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/files")
			So(stub.GetParams.Get("account_id"), ShouldEqual, "21")

			So(len(files), ShouldEqual, 1)

			v := &TestValue{}
			v.Count = 1
			CheckFile(v, files[0])
		})

		Convey("GetFilesRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.GetFilesRaw(42, 21)
			So(string(b), ShouldEqual, correctJSON)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/files")
			So(stub.GetParams.Get("account_id"), ShouldEqual, "21")
		})
	})
}

func TestGetSpecificFilesRaw(t *testing.T) {
	testToken := "testToken"
	client := New(testToken)

	Convey("correct", t, func() {
		correctJSON := `
{
  "file_id": 1,
  "account": {
    "account_id": 2,
    "name": "string_3",
    "avatar_image_url": "string_4"
  },
  "message_id": 5,
  "filename": "string_6",
  "filesize": 7,
  "upload_time": 8,
  "download_url": "string_9"
}
`
		Convey("GetSpecificFiles", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			file, err := client.GetSpecificFiles(42, 21, true)
			So(err, ShouldBeNil)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/files/21")
			So(stub.GetParams.Get("create_download_url"), ShouldEqual, "1")

			v := &TestValue{}
			v.Count = 1
			CheckFile(v, file)
			So(file.DownloadURL, ShouldEqual, v.GetString())
		})

		Convey("GetSpecificFilesRaw", func() {
			stub := &stubHTTP{}
			stub.GetByte = []byte(correctJSON)
			client.connection = stub

			b, _ := client.GetSpecificFilesRaw(42, 21, true)
			So(string(b), ShouldEqual, correctJSON)
			So(stub.GetCount, ShouldEqual, 1)
			So(stub.GetEndPoint, ShouldEqual, "rooms/42/files/21")
			So(stub.GetParams.Get("create_download_url"), ShouldEqual, "1")
		})
	})
}
