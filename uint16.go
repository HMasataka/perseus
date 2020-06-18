package perseus

// InUint16 returns whether i is in list
func InUint16(i uint16, list []uint16) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexUint16 returns the position of s in list. If s is not found, return -1.
func IndexUint16(s uint16, list []uint16) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftUint16 returns the first element of slice and other element's slice.
func ShiftUint16(slice []uint16) (uint16, []uint16) {
	return slice[0], slice[1:]
}

// UnshiftUint16 add an element to the beginning of a slice.
func UnshiftUint16(sep uint16, i []uint16) []uint16 {
	return append([]uint16{sep}, i...)
}

// DeleteUint16 delete specified element from slice
func DeleteUint16(slice []uint16, sep uint16) []uint16 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutUint16 Delete from i to j from the slice
func CutUint16(slice []uint16, i, j uint16) []uint16 {
	return append(slice[:i], slice[j:]...)
}

// InsertUint16 insert element to specified position
func InsertUint16(slice []uint16, element, position uint16) []uint16 {
	return append(slice[:position], append([]uint16{element}, slice[position:]...)...)
}

// InsertVectorUint16 insert slice to specified position
func InsertVectorUint16(origin, insert []uint16, position uint16) []uint16 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopUint16 returns the last element of slice and other element's slice.
func PopUint16(slice []uint16) (uint16, []uint16) {
	return slice[0], slice[1:]
}

// ReversedUint16 returns reversed slice
func ReversedUint16(slice []uint16) []uint16 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendUint16 connect slices together
func ExtendUint16(A, B []uint16) []uint16 {
	return append(A, B...)
}

func sumUint16(values ...uint16) uint64 {
	var sum uint64
	for _, v := range values {
		sum += uint64(v)
	}
	return sum
}

// SumUint16 calculate summaries of arguments
func SumUint16(values ...uint16) uint16 {
	return uint16(sumUint16(values...))
}

// SumUint16ToUint64 calculate summaries of arguments
func SumUint16ToUint64(values ...uint16) uint64 {
	return sumUint16(values...)
}

// SumUint16ToUint calculate summaries of arguments
func SumUint16ToUint(values ...uint16) uint {
	return uint(sumUint16(values...))
}
