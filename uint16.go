package perseus

func InUint16(i uint16, list []uint16) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexUint16(s uint16, list []uint16) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftUint16(slice []uint16) (uint16, []uint16) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftUint16(sep uint16, i []uint16) []uint16 {
	return append([]uint16{sep}, i...)
}

func DeleteUint16(slice []uint16, sep uint16) []uint16 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutUint16(slice []uint16, i, j uint16) []uint16 {
	return append(slice[:i], slice[j:]...)
}

func InsertUint16(slice []uint16, element, position uint16) []uint16 {
	return append(slice[:position], append([]uint16{element}, slice[position:]...)...)
}

func InsertVectorUint16(origin, insert []uint16, position uint16) []uint16 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopUint16(slice []uint16) (uint16, []uint16) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedUint16(slice []uint16) []uint16 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendUint16(A, B []uint16) []uint16 {
	return append(A, B...)
}
