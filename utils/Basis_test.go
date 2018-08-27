/*
Description:
test Basis.go method
Cmd: go test -v Basis.go Basis_test.go
 * Author: architect.bian
 * Date: 2018/08/27 13:59
 */
package utils

import (
	"testing"
)

func TestBool(t *testing.T) {
	var b bool
	var bp *bool
	if b != false {
		t.Error("b bool default value should be false")
	}
	if bp != nil {
		t.Error("bp *bool default value should be nil")
	}
	bp = nil
	bp = Basis.Bool(true)
	if bp == nil {
		t.Error("bp *bool current value should not be nil")
	}
	if *bp != true {
		t.Error("bp *bool current value should be true")
	}
}

func TestInt(t *testing.T) {
	var v int
	var vp *int
	if v != 0 {
		t.Error("v default value should be 0")
	}
	if vp != nil {
		t.Error("v default value should be nil")
	}
	vp = nil
	vp = Basis.Int(-1)
	if vp == nil {
		t.Error("vp current value should not be nil")
	}
	if *vp != -1 {
		t.Error("vp current value should be -1")
	}
}

func TestString(t *testing.T) {
	var v string
	var vp *string
	if v != "" {
		t.Error("v default value should be empty")
	}
	if vp != nil {
		t.Error("v default value should be nil")
	}
	vp = nil
	vp = Basis.String("s")
	if vp == nil {
		t.Error("vp current value should not be nil")
	}
	if *vp != "s" {
		t.Error("vp current value should be s")
	}
}