package main

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestTwoSum(t *testing.T) {
	Convey("start TwoSum [2, 7, 11, 15], target=9,", t, func() {
		r := TwoSum([]int{2, 7, 11, 15}, 9)
		Convey("The value should be equal", func() {
			So(r[0], ShouldEqual, 0)
			So(r[1], ShouldEqual, 1)
		})
	})

}
