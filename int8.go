package perseus

func InInt8(i int8, list []int8) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexInt8(s int8, list []int8) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftInt8(slice []int8) (int8, []int8) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftInt8(sep int8, i []int8) []int8 {
	return append([]int8{sep}, i...)
}

func DeleteInt8(slice []int8, sep int8) []int8 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutInt8(slice []int8, i, j int8) []int8 {
	return append(slice[:i], slice[j:]...)
}

func InsertInt8(slice []int8, element, position int8) []int8 {
	return append(slice[:position], append([]int8{element}, slice[position:]...)...)
}

func InsertVectorInt8(origin, insert []int8, position int8) []int8 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopInt8(slice []int8) (int8, []int8) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedInt8(slice []int8) []int8 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendInt8(A, B []int8) []int8 {
	return append(A, B...)
}
