package perseus

// InInt8 returns whether i is in list
func InInt8(i int8, list []int8) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexInt8 returns the position of s in list. If s is not found, return -1.
func IndexInt8(s int8, list []int8) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftInt8 returns the first element of slice and other element's slice.
func ShiftInt8(slice []int8) (int8, []int8) {
	return slice[0], slice[1:]
}

// UnshiftInt8 add an element to the beginning of a slice.
func UnshiftInt8(sep int8, i []int8) []int8 {
	return append([]int8{sep}, i...)
}

// DeleteInt8 delete specified element from slice
func DeleteInt8(slice []int8, sep int8) []int8 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutInt8 Delete from i to j from the slice
func CutInt8(slice []int8, i, j int8) []int8 {
	return append(slice[:i], slice[j:]...)
}

// InsertInt8 insert element to specified position
func InsertInt8(slice []int8, element, position int8) []int8 {
	return append(slice[:position], append([]int8{element}, slice[position:]...)...)
}

// InsertVectorInt8 insert slice to specified position
func InsertVectorInt8(origin, insert []int8, position int8) []int8 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopInt8 returns the last element of slice and other element's slice.
func PopInt8(slice []int8) (int8, []int8) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// ReversedInt8 returns reversed slice
func ReversedInt8(slice []int8) []int8 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendInt8 connect slices together
func ExtendInt8(A, B []int8) []int8 {
	return append(A, B...)
}
