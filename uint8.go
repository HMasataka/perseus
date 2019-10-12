package perseus

func InUint8(i uint8, list []uint8) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexUint8(s uint8, list []uint8) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftUint8(slice []uint8) (uint8, []uint8) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftUint8(sep uint8, i []uint8) []uint8 {
	return append([]uint8{sep}, i...)
}

func DeleteUint8(slice []uint8, sep uint8) []uint8 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutUint8(slice []uint8, i, j uint8) []uint8 {
	return append(slice[:i], slice[j:]...)
}

func InsertUint8(slice []uint8, element, position uint8) []uint8 {
	return append(slice[:position], append([]uint8{element}, slice[position:]...)...)
}

func InsertVectorUint8(origin, insert []uint8, position uint8) []uint8 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopUint8(slice []uint8) (uint8, []uint8) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedUint8(slice []uint8) []uint8 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendUint8(A, B []uint8) []uint8 {
	return append(A, B...)
}
