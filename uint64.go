package perseus

func InUint64(i uint64, list []uint64) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexUint64(s uint64, list []uint64) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftUint64(slice []uint64) (uint64, []uint64) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftUint64(sep uint64, i []uint64) []uint64 {
	return append([]uint64{sep}, i...)
}

func DeleteUint64(slice []uint64, sep uint64) []uint64 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutUint64(slice []uint64, i, j uint64) []uint64 {
	return append(slice[:i], slice[j:]...)
}

func InsertUint64(slice []uint64, element, position uint64) []uint64 {
	return append(slice[:position], append([]uint64{element}, slice[position:]...)...)
}

func InsertVectorUint64(origin, insert []uint64, position uint64) []uint64 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopUint64(slice []uint64) (uint64, []uint64) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedUint64(slice []uint64) []uint64 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendUint64(A, B []uint64) []uint64 {
	return append(A, B...)
}
