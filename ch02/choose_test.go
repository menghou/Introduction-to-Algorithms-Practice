package ch02

import "testing"

func TestChoose(t *testing.T) {
	a := []int{5, 2, 6, 1}
	b := []int{1, 2, 5, 6}
	Choose(a)
	for index, value := range a {
		if value != b[index] {
			t.Errorf("index: %d, a: %d, b: %d", index, value, b[index])
		}
	}
}
