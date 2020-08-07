package main

import (
	"testing"
)

func TestFun(t *testing.T) {

	logger := Constructor()

	v := logger.ShouldPrintMessage(1, "foo")
	if v != true {
		t.Errorf("got %t, want %t", v, true)
	}
	v = logger.ShouldPrintMessage(2, "bar")
	if v != true {
		t.Errorf("got %t, want %t", v, true)
	}
	v = logger.ShouldPrintMessage(3, "foo")
	if v != false {
		t.Errorf("got %t, want %t", v, false)
	}
	v = logger.ShouldPrintMessage(8, "bar")
	if v != false {
		t.Errorf("got %t, want %t", v, false)
	}
	v = logger.ShouldPrintMessage(10, "foo")
	if v != false {
		t.Errorf("got %t, want %t", v, false)
	}
	v = logger.ShouldPrintMessage(11, "foo")
	if v != true {
		t.Errorf("got %t, want %t", v, true)
	}

}
