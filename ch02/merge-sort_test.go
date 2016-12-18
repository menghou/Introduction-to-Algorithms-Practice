package ch02

import "testing"

func TestMerge(t *testing.T) {
	a := []int{2, 7, 3, 4, 9, 12}
	var p, q, r int = 0, 1, 5
	Merge(a, p, q, r)
	b := []int{2, 3, 4, 7, 9, 12}
	for index, value := range a {
		if value != b[index] {
			t.Errorf("index: %d, a:%d, b:%d", index, value, b[index])
		}
	}
}

func TestMergeSort(t *testing.T) {
	a := []int{5, 2, 4, 7, 1, 3, 2, 9}
	b := []int{1, 2, 2, 3, 4, 5, 7, 9}
	MergeSort(a, 0, 7)
	for index, value := range a {
		if value != b[index] {
			t.Errorf("index: %d, a:%d, b:%d", index, value, b[index])
		}
	}
}
