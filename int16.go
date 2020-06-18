package perseus

// InInt16 returns whether i is in list
func InInt16(i int16, list []int16) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexInt16 returns the position of s in list. If s is not found, return -1.
func IndexInt16(s int16, list []int16) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftInt16 returns the first element of slice and other element's slice.
func ShiftInt16(slice []int16) (int16, []int16) {
	return slice[0], slice[1:]
}

// UnshiftInt16 add an element to the beginning of a slice.
func UnshiftInt16(sep int16, i []int16) []int16 {
	return append([]int16{sep}, i...)
}

// DeleteInt16 delete specified element from slice
func DeleteInt16(slice []int16, sep int16) []int16 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutInt16 Delete from i to j from the slice
func CutInt16(slice []int16, i, j int16) []int16 {
	return append(slice[:i], slice[j:]...)
}

// InsertInt16 insert element to specified position
func InsertInt16(slice []int16, element, position int16) []int16 {
	return append(slice[:position], append([]int16{element}, slice[position:]...)...)
}

// InsertVectorInt16 insert slice to specified position
func InsertVectorInt16(origin, insert []int16, position int16) []int16 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopInt16 returns the last element of slice and other element's slice.
func PopInt16(slice []int16) (int16, []int16) {
	return slice[0], slice[1:]
}

// ReversedInt16 returns reversed slice
func ReversedInt16(slice []int16) []int16 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendInt16 connect slices together
func ExtendInt16(A, B []int16) []int16 {
	return append(A, B...)
}

// SumInt16 calculate summaries of arguments
func SumInt16(values ...int16) int16 {
	var sum int16
	for _, v := range values {
		sum += v
	}
	return sum
}

// SumInt16ToInt64 calculate summaries of arguments
func SumInt16ToInt64(values ...int16) int64 {
	var sum int64
	for _, v := range values {
		sum += int64(v)
	}
	return sum
}
