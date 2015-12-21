package openstack

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConnection(t *testing.T) {
	var sut Connection
	Convey("Connection should have options", t, func() {
		So(sut.connect(), ShouldEqual, true)
	})
}
