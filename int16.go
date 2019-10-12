package perseus

func InInt16(i int16, list []int16) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexInt16(s int16, list []int16) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftInt16(slice []int16) (int16, []int16) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftInt16(sep int16, i []int16) []int16 {
	return append([]int16{sep}, i...)
}

func DeleteInt16(slice []int16, sep int16) []int16 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutInt16(slice []int16, i, j int16) []int16 {
	return append(slice[:i], slice[j:]...)
}

func InsertInt16(slice []int16, element, position int16) []int16 {
	return append(slice[:position], append([]int16{element}, slice[position:]...)...)
}

func InsertVectorInt16(origin, insert []int16, position int16) []int16 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopInt16(slice []int16) (int16, []int16) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedInt16(slice []int16) []int16 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendInt16(A, B []int16) []int16 {
	return append(A, B...)
}
