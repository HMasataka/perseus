package perseus

// InInt32 returns whether i is in list
func InInt32(i int32, list []int32) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexInt32 returns the position of s in list. If s is not found, return -1.
func IndexInt32(s int32, list []int32) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftInt32 returns the first element of slice and other element's slice.
func ShiftInt32(slice []int32) (int32, []int32) {
	return slice[0], slice[1:]
}

// UnshiftInt32 add an element to the beginning of a slice.
func UnshiftInt32(sep int32, i []int32) []int32 {
	return append([]int32{sep}, i...)
}

// DeleteInt32 delete specified element from slice
func DeleteInt32(slice []int32, sep int32) []int32 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutInt32 Delete from i to j from the slice
func CutInt32(slice []int32, i, j int32) []int32 {
	return append(slice[:i], slice[j:]...)
}

// InsertInt32 insert element to specified position
func InsertInt32(slice []int32, element, position int32) []int32 {
	return append(slice[:position], append([]int32{element}, slice[position:]...)...)
}

// InsertVectorInt32 insert slice to specified position
func InsertVectorInt32(origin, insert []int32, position int32) []int32 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopInt32 returns the last element of slice and other element's slice.
func PopInt32(slice []int32) (int32, []int32) {
	return slice[0], slice[1:]
}

// ReversedInt32 returns reversed slice
func ReversedInt32(slice []int32) []int32 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendInt32 connect slices together
func ExtendInt32(A, B []int32) []int32 {
	return append(A, B...)
}

func sumInt32(values ...int32) int64 {
	var sum int64
	for _, v := range values {
		sum += int64(v)
	}
	return sum
}

// SumInt32 calculate summaries of arguments
func SumInt32(values ...int32) int32 {
	return int32(sumInt32(values...))
}

// SumInt32ToInt calculate summaries of arguments
func SumInt32ToInt(values ...int32) int {
	return int(sumInt32(values...))
}

// SumInt32ToInt64 calculate summaries of arguments
func SumInt32ToInt64(values ...int32) int64 {
	return sumInt32(values...)
}
