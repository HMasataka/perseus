package perseus

func InUint(i uint, list []uint) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexUint(s uint, list []uint) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftUint(slice []uint) (uint, []uint) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftUint(sep uint, i []uint) []uint {
	return append([]uint{sep}, i...)
}

func DeleteUint(slice []uint, sep uint) []uint {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutUint(slice []uint, i, j uint) []uint {
	return append(slice[:i], slice[j:]...)
}

func InsertUint(slice []uint, element, position uint) []uint {
	return append(slice[:position], append([]uint{element}, slice[position:]...)...)
}

func InsertVectorUint(origin, insert []uint, position uint) []uint {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopUint(slice []uint) (uint, []uint) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedUint(slice []uint) []uint {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendUint(A, B []uint) []uint {
	return append(A, B...)
}
