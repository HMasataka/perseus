package perseus

// InInt64 returns whether i is in list
func InInt64(i int64, list []int64) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexInt64 returns the position of s in list. If s is not found, return -1.
func IndexInt64(s int64, list []int64) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftInt64 returns the first element of slice and other element's slice.
func ShiftInt64(slice []int64) (int64, []int64) {
	x, s := slice[0], slice[1:]
	return x, s
}

// UnshiftInt64 add an element to the beginning of a slice.
func UnshiftInt64(sep int64, i []int64) []int64 {
	return append([]int64{sep}, i...)
}

// DeleteInt64 delete specified element from slice
func DeleteInt64(slice []int64, sep int64) []int64 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutInt64 Delete from i to j from the slice
func CutInt64(slice []int64, i, j int64) []int64 {
	return append(slice[:i], slice[j:]...)
}

// InsertInt64 insert element to specified position
func InsertInt64(slice []int64, element, position int64) []int64 {
	return append(slice[:position], append([]int64{element}, slice[position:]...)...)
}

// InsertVectorInt64 insert slice to specified position
func InsertVectorInt64(origin, insert []int64, position int64) []int64 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopInt64 returns the last element of slice and other element's slice.
func PopInt64(slice []int64) (int64, []int64) {
	return slice[0], slice[1:]
}

// ReversedInt64 returns reversed slice
func ReversedInt64(slice []int64) []int64 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendInt64 connect slices together
func ExtendInt64(A, B []int64) []int64 {
	return append(A, B...)
}

func sumInt64(values ...int64) int64 {
	var sum int64
	for _, v := range values {
		sum += v
	}
	return sum
}

// SumInt64 calculate summaries of arguments
func SumInt64(values ...int64) int64 {
	return sumInt64(values...)
}

// SumInt64ToInt calculate summaries of arguments
func SumInt64ToInt(values ...int64) int {
	return int(sumInt64(values...))
}
