package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionGet(t *testing.T) {
	var req http.Request
	res := httptest.NewRecorder()
	VersionGet(res, &req)
	Convey("The header should be application/json", t, func() {
		So(res.HeaderMap.Get("Content-Type"), ShouldEqual, "application/json")
	})
	Convey("The body should contain expected json", t, func() {
		So(res.Body.String(), ShouldEqual, `{"versions": [{"status": "stable", "updated": "2015-12-11T00:00:00Z", "id": "v1"}]}`)
	})
}
