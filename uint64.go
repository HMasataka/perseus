package perseus

// InUint64 returns whether i is in list
func InUint64(i uint64, list []uint64) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexUint64 returns the position of s in list. If s is not found, return -1.
func IndexUint64(s uint64, list []uint64) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftUint64 returns the first element of slice and other element's slice.
func ShiftUint64(slice []uint64) (uint64, []uint64) {
	return slice[0], slice[1:]
}

// UnshiftUint64 add an element to the beginning of a slice.
func UnshiftUint64(sep uint64, i []uint64) []uint64 {
	return append([]uint64{sep}, i...)
}

// DeleteUint64 delete specified element from slice
func DeleteUint64(slice []uint64, sep uint64) []uint64 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutUint64 Delete from i to j from the slice
func CutUint64(slice []uint64, i, j uint64) []uint64 {
	return append(slice[:i], slice[j:]...)
}

// InsertUint64 insert element to specified position
func InsertUint64(slice []uint64, element, position uint64) []uint64 {
	return append(slice[:position], append([]uint64{element}, slice[position:]...)...)
}

// InsertVectorUint64 insert slice to specified position
func InsertVectorUint64(origin, insert []uint64, position uint64) []uint64 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopUint64 returns the last element of slice and other element's slice.
func PopUint64(slice []uint64) (uint64, []uint64) {
	return slice[0], slice[1:]
}

// ReversedUint64 returns reversed slice
func ReversedUint64(slice []uint64) []uint64 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendUint64 connect slices together
func ExtendUint64(A, B []uint64) []uint64 {
	return append(A, B...)
}

// SumUint64 calculate summaries of arguments
func SumUint64(values ...uint64) uint64 {
	var sum uint64
	for _, v := range values {
		sum += v
	}
	return sum
}
