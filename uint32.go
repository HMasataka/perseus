package perseus

// InUint32 returns whether i is in list
func InUint32(i uint32, list []uint32) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexUint32 returns the position of s in list. If s is not found, return -1.
func IndexUint32(s uint32, list []uint32) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftUint32 returns the first element of slice and other element's slice.
func ShiftUint32(slice []uint32) (uint32, []uint32) {
	return slice[0], slice[1:]
}

// UnshiftUint32 add an element to the beginning of a slice.
func UnshiftUint32(sep uint32, i []uint32) []uint32 {
	return append([]uint32{sep}, i...)
}

// DeleteUint32 delete specified element from slice
func DeleteUint32(slice []uint32, sep uint32) []uint32 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutUint32 Delete from i to j from the slice
func CutUint32(slice []uint32, i, j uint32) []uint32 {
	return append(slice[:i], slice[j:]...)
}

// InsertUint32 element to specified position
func InsertUint32(slice []uint32, element, position uint32) []uint32 {
	return append(slice[:position], append([]uint32{element}, slice[position:]...)...)
}

// InsertVectorUint32 slice to specified position
func InsertVectorUint32(origin, insert []uint32, position uint32) []uint32 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopUint32 returns the last element of slice and other element's slice.
func PopUint32(slice []uint32) (uint32, []uint32) {
	return slice[0], slice[1:]
}

// ReversedUint32 returns reversed slice
func ReversedUint32(slice []uint32) []uint32 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendUint32 connect slices together
func ExtendUint32(A, B []uint32) []uint32 {
	return append(A, B...)
}
