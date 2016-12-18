package ch02

const (
	NIL = -1
)

func LinearSearch(a []int, b int) int {
	var indexOutPut int = NIL
	for index, item := range a {
		if item == b {
			indexOutPut = index
			break
		}
	}
	return indexOutPut
}

func BinarySearch(a []int, b int) int {
	high := len(a) - 1
	low := 0
	for {
		if high == low {
			if a[high] != b {
				return NIL
			} else {
				return high
			}
		}
		if high-low == 1 {
			if a[high] == b {
				return high
			} else if a[low] == b {
				return low
			} else {
				return NIL
			}
		}
		midIndex := (high + low) / 2
		if a[midIndex] < b {
			low = midIndex
		} else if a[midIndex] > b {
			high = midIndex
		} else {
			return midIndex
		}
	}

}
