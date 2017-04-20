package main

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestReverseInteger(t *testing.T) {
	Convey("Test TestReverseInteger", t, func() {
		Convey("Test Neg", func() {
			So(321, ShouldEqual, Reverse(123))
			So(-321, ShouldEqual, Reverse(-123))
			So(0, ShouldEqual, Reverse(1534236469))
		})
	})

}
