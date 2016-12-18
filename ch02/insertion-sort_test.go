package ch02

import "testing"

func TestInsertionSortInt(t *testing.T) {
	a := []int{5, 2, 6, 1}
	InsertionSortInt(a)
	b := []int{1, 2, 5, 6}
	for index, _ := range a {
		if a[index] != b[index] {
			t.Errorf("index: %d, a:%d, b:%d", index, a[index], b[index])
		}
	}
}

func TestInsertionSortIntReverse(t *testing.T) {
	a := []int{5, 2, 6, 1}
	c := []int{6, 5, 2, 1}
	InsertionSortIntReverse(a)
	for index, _ := range a {
		if a[index] != c[index] {
			t.Errorf("index: %d, a:%d, b:%d", index, a[index], c[index])
		}
	}
}
func TestInsertionSortRecursion(t *testing.T) {
	a := []int{5, 2, 6, 1}
	InsertionSortRecursion(a, 4)
	b := []int{1, 2, 5, 6}
	for index, _ := range a {
		if a[index] != b[index] {
			t.Errorf("index: %d, a:%d, b:%d", index, a[index], b[index])
		}
	}
}
func TestInsertionsortRecursion2(t *testing.T) {
	a := []int{2, 5, 6, 1}
	insertionSortRecursion(a, 3)
	b := []int{1, 2, 5, 6}
	for index, _ := range a {
		if a[index] != b[index] {
			t.Errorf("index: %d, a:%d, b:%d", index, a[index], b[index])
		}
	}
}
