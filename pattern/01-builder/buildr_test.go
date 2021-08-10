package _1_builder

import "testing"

func testBuilderPattern(t *testing.T) {
	txtBuilder := TxtBuilder{}
	director := Director{}
	director.SetBuilder(txtBuilder)

}
