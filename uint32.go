package perseus

func InUint32(i uint32, list []uint32) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexUint32(s uint32, list []uint32) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftUint32(slice []uint32) (uint32, []uint32) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftUint32(sep uint32, i []uint32) []uint32 {
	return append([]uint32{sep}, i...)
}

func DeleteUint32(slice []uint32, sep uint32) []uint32 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutUint32(slice []uint32, i, j uint32) []uint32 {
	return append(slice[:i], slice[j:]...)
}

func InsertUint32(slice []uint32, element, position uint32) []uint32 {
	return append(slice[:position], append([]uint32{element}, slice[position:]...)...)
}

func InsertVectorUint32(origin, insert []uint32, position uint32) []uint32 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopUint32(slice []uint32) (uint32, []uint32) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedUint32(slice []uint32) []uint32 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendUint32(A, B []uint32) []uint32 {
	return append(A, B...)
}
