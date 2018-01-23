package vals

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func withArray() map[string]interface{} {
	m := make(map[string]interface{})
	m["tags"] = []string{
		"cafe",
		"coffee house",
	}
	return m
}

func singleLevelMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["city"] = "Boulder"
	m["firstName"] = "Bruce"
	m["lastName"] = "Wayne"
	return m
}

func TestValue(t *testing.T) {

	Convey("With array ", t, func() {
		m := withArray()
		val := New(m).At("tags")

		So(val, ShouldNotBeNil)
		So(val.data, ShouldNotBeNil)

		fmt.Println(val.In(0).AsString())
		So(val.In(0).AsString(), ShouldEqual, "cafe")
		So(val.In(1).AsString(), ShouldEqual, "coffee house")
	})

	Convey("Value ", t, func() {
		m := singleLevelMap()
		val := New(m)

		for k, v := range m {
			So(val.At(k).AsString(), ShouldEqual, v)
		}
	})

	Convey("Initial test", t, func() {
		fmt.Println("Initial Test")
		So(true, ShouldEqual, true)
	})
}
