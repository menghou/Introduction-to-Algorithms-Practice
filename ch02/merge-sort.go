package ch02

func MergeSort(a []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(a, p, q)
		MergeSort(a, q+1, r)
		Merge(a, p, q, r)
	}
}

func Merge(a []int, p, q, r int) {
	var i int = 0
	var j int = 0
	var k int = p
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 0; i < q+1-p; i++ {
		left = append(left, a[i+p])
	}
	for i := 0; i < r-q; i++ {
		right = append(right, a[i+q+1])
	}
	for i != (q-p+1) && j != (r-q) {
		if left[i] < right[j] {
			a[k] = left[i]
			i += 1
			k += 1
		} else {
			a[k] = right[j]
			j += 1
			k += 1
		}
	}
	if i == q-p+1 {
		for k != r+1 && j != (r-q) {
			a[k] = right[j]
			k += 1
			j += 1
		}
	} else {
		for k != r+1 && i != (p-q+1) {
			a[k] = left[i]
			k += 1
			i += 1
		}
	}
}
