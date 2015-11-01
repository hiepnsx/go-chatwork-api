package gochatwork

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	testToken := "testToken"

	Convey("normal", t, func() {
		client := New(testToken)
		So(client.config.url, ShouldEqual, chatworkURL)
		So(client.config.token, ShouldEqual, testToken)
	})
}

type TestValue struct {
	Count int64
}

func (t *TestValue) getCount() int64 {
	c := t.Count
	t.Count++
	return c
}

func (t *TestValue) GetInt64() int64 {
	return t.getCount()
}

func (t *TestValue) GetString() string {
	return "string_" + strconv.FormatInt(t.getCount(), 10)
}

type stubHTTP struct {
	GetCount    int
	GetByte     []byte
	GetEndPoint string
	GetParams   url.Values

	PostCount    int
	PostByte     []byte
	PostEndPoint string
	PostParams   url.Values

	PutCount    int
	PutByte     []byte
	PutEndPoint string
	PutParams   url.Values

	DeleteCount    int
	DeleteByte     []byte
	DeleteEndPoint string
	DeleteParams   url.Values
}

func (h *stubHTTP) Get(endPoint string, params url.Values, config *config) ([]byte, error) {
	h.GetCount++
	h.GetEndPoint = endPoint
	h.GetParams = params
	return h.GetByte, nil
}

func (h *stubHTTP) Post(endPoint string, params url.Values, config *config) ([]byte, error) {
	h.PostCount++
	h.PostEndPoint = endPoint
	h.PostParams = params
	return h.PostByte, nil
}

func (h *stubHTTP) Put(endPoint string, params url.Values, config *config) ([]byte, error) {
	h.PutCount++
	h.PutEndPoint = endPoint
	h.PutParams = params
	return h.PutByte, nil
}

func (h *stubHTTP) Delete(endPoint string, params url.Values, config *config) ([]byte, error) {
	h.DeleteCount++
	h.DeleteEndPoint = endPoint
	h.DeleteParams = params
	return h.DeleteByte, nil
}
