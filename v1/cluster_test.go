package v1

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClusterGet(t *testing.T) {
	var req http.Request
	res := httptest.NewRecorder()
	ClusterGet(res, &req)
	Convey("The header should be application/json", t, func() {
		So(res.HeaderMap.Get("Content-Type"), ShouldEqual, "application/json")
	})
	Convey("The body should contain expected json", t, func() {
		So(res.Body.String(), ShouldEqual, `{"clusters": []}`)
	})
}
