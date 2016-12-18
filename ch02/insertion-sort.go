package ch02

//this file implements InsertionSort function described in ch02

func InsertionSortInt(A []int) {
	length := len(A)
	if length <= 1 {
		return
	} else {
		for i := 1; i <= length-1; i++ {
			key := A[i]
			var j int = i - 1
			for j >= 0 && key < A[j] {
				A[j+1] = A[j]
				j -= 1
			}
			A[j+1] = key
		}
	}
}

func InsertionSortIntReverse(A []int) {
	length := len(A)
	if length <= 1 {
		return
	} else {
		for i := 1; i <= length-1; i++ {
			key := A[i]
			var j int = i - 1
			for j >= 0 && key > A[j] {
				A[j+1] = A[j]
				j -= 1
			}
			A[j+1] = key
		}
	}
}

func InsertionSortRecursion(a []int, n int) {
	for n > 0 {
		InsertionSortRecursion(a, n-1)
		insertionSortRecursion(a, n)
	}
}

func insertionSortRecursion(a []int, n int) {
	key := a[n]
	var j int = n - 1
	for j >= 0 && key < a[j] {
		a[j+1] = a[j]
		j -= 1
	}
	a[j+1] = key
}
