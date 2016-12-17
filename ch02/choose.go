package ch02

func Choose(a []int) {
	length := len(a)
	if length <= 1 {
		return
	}
	for i, item := range a {
		minimal := item
		minimalIndex := i
		for j := i + 1; j <= len(a)-1; j++ {
			if a[j] < minimal {
				minimal = a[j]
				minimalIndex = j
			}
		}
		if minimal < item {
			a[i] ^= a[minimalIndex]
			a[minimalIndex] ^= a[i]
			a[i] ^= a[minimalIndex]
		}
	}
}
