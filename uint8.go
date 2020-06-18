package perseus

// InUint8 returns whether i is in list
func InUint8(i uint8, list []uint8) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexUint8 returns the position of s in list. If s is not found, return -1.
func IndexUint8(s uint8, list []uint8) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftUint8 returns the first element of slice and other element's slice.
func ShiftUint8(slice []uint8) (uint8, []uint8) {
	return slice[0], slice[1:]
}

// UnshiftUint8 add an element to the beginning of a slice.
func UnshiftUint8(sep uint8, i []uint8) []uint8 {
	return append([]uint8{sep}, i...)
}

// DeleteUint8 specified element from slice
func DeleteUint8(slice []uint8, sep uint8) []uint8 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutUint8 Delete from i to j from the slice
func CutUint8(slice []uint8, i, j uint8) []uint8 {
	return append(slice[:i], slice[j:]...)
}

// InsertUint8 element to specified position
func InsertUint8(slice []uint8, element, position uint8) []uint8 {
	return append(slice[:position], append([]uint8{element}, slice[position:]...)...)
}

// InsertVectorUint8 slice to specified position
func InsertVectorUint8(origin, insert []uint8, position uint8) []uint8 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopUint8 returns the last element of slice and other element's slice.
func PopUint8(slice []uint8) (uint8, []uint8) {
	return slice[0], slice[1:]
}

// ReversedUint8 returns reversed slice
func ReversedUint8(slice []uint8) []uint8 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendUint8 connect slices together
func ExtendUint8(A, B []uint8) []uint8 {
	return append(A, B...)
}

// SumUint8 calculate summaries of arguments
func SumUint8(values ...uint8) uint8 {
	var sum uint8
	for _, v := range values {
		sum += v
	}
	return sum
}

// SumUint8ToUint64 calculate summaries of arguments
func SumUnt8ToUint64(values ...uint) uint64 {
	var sum uint64
	for _, v := range values {
		sum += uint64(v)
	}
	return sum
}
