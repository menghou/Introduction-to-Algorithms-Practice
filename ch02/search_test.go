package ch02

import "testing"

var a []int = []int{5, 7, 9}
var b int = 5
var c int = 11

func TestLinearSearch(t *testing.T) {
	indexb := LinearSearch(a, b)
	if indexb != 0 {
		t.Errorf("indexb :%d", indexb)
	}
	indexc := LinearSearch(a, c)
	if indexc != NIL {
		t.Errorf("indexc :%d", indexc)
	}
}

func TestBinarySearch(t *testing.T) {
	indexb := BinarySearch(a, b)
	if indexb != 0 {
		t.Errorf("indexb :%d", indexb)
	}
	indexc := BinarySearch(a, c)
	if indexc != NIL {
		t.Errorf("indexc :%d", indexc)
	}
}
