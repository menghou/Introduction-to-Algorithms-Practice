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
