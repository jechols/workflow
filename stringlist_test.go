package main

import (
	"testing"
	"strings"
)

func assertEqualS(expected, actual string, message string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %#v, but got %#v - %s", expected, actual, message)
	}
}

func assertEqualI(expected, actual int, message string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %#v, but got %#v - %s", expected, actual, message)
	}
}

func TestStuff(t *testing.T) {
	slist := NewStringList()
	slist.AddString("string 1")
	slist.AddString("string 2")
	slist.AddString("string 3")

	source := strings.Join(slist.source.data, ",")
	assertEqualS("string 1,string 2,string 3", source, "Stringlist value", t)

	assertEqualI(3, slist.source.Len(), "Source list size should be 3", t)
	assertEqualI(0, slist.options.Len(), "Options should be empty at first", t)

	_ = slist.RandomString()
	assertEqualI(3, slist.source.Len(), "Source list size should still be 3", t)
	assertEqualI(2, slist.options.Len(), "Options should include the remaining two strings", t)
}
