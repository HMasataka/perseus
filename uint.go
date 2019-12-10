package perseus

// InUint returns whether i is in list
func InUint(i uint, list []uint) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexUint returns the position of s in list. If s is not found, return -1.
func IndexUint(s uint, list []uint) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftUint returns the first element of slice and other element's slice.
func ShiftUint(slice []uint) (uint, []uint) {
	return slice[0], slice[1:]
}

// UnshiftUint add an element to the beginning of a slice.
func UnshiftUint(sep uint, i []uint) []uint {
	return append([]uint{sep}, i...)
}

// DeleteUint delete specified element from slice
func DeleteUint(slice []uint, sep uint) []uint {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutUint Delete from i to j from the slice
func CutUint(slice []uint, i, j uint) []uint {
	return append(slice[:i], slice[j:]...)
}

// InsertUint insert element to specified position
func InsertUint(slice []uint, element, position uint) []uint {
	return append(slice[:position], append([]uint{element}, slice[position:]...)...)
}

// InsertVectorUint insert slice to specified position
func InsertVectorUint(origin, insert []uint, position uint) []uint {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopUint returns the last element of slice and other element's slice.
func PopUint(slice []uint) (uint, []uint) {
	return slice[0], slice[1:]
}

// ReversedUint returns reversed slice
func ReversedUint(slice []uint) []uint {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendUint connect slices together
func ExtendUint(A, B []uint) []uint {
	return append(A, B...)
}
