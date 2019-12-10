package perseus

// InInt returns whether i is in list
func InInt(i int, list []int) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexInt returns the position of s in list. If s is not found, return -1.
func IndexInt(s int, list []int) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftInt returns the first element of slice and other element's slice.
func ShiftInt(slice []int) (int, []int) {
	return slice[0], slice[1:]
}

// UnshiftInt add an element to the beginning of a slice.
func UnshiftInt(sep int, i []int) []int {
	return append([]int{sep}, i...)
}

// DeleteInt Delete specified element from slice
func DeleteInt(slice []int, sep int) []int {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutInt Delete from i to j from the slice
func CutInt(slice []int, i, j int) []int {
	return append(slice[:i], slice[j:]...)
}

// InsertInt Insert element to specified position
func InsertInt(slice []int, element, position int) []int {
	return append(slice[:position], append([]int{element}, slice[position:]...)...)
}

// InsertVectorInt Insert slice to specified position
func InsertVectorInt(origin, insert []int, position int) []int {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopInt returns the last element of slice and other element's slice.
func PopInt(slice []int) (int, []int) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// ReversedInt returns reversed slice
func ReversedInt(slice []int) []int {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendInt connect slices together
func ExtendInt(A, B []int) []int {
	return append(A, B...)
}
