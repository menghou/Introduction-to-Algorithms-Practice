package ch02

import "testing"

const (
	a1 = 100
	b1 = 3
	c1 = 103
)

func TestBinarySum(t *testing.T) {
	if BinarySum(a1, b1) != c1 {
		t.Errorf("output :%d", BinarySum(a1, b1))
	}
}
