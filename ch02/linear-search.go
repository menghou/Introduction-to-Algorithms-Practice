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
