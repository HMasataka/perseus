package perseus

func InInt32(i int32, list []int32) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexInt32(s int32, list []int32) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftInt32(slice []int32) (int32, []int32) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftInt32(sep int32, i []int32) []int32 {
	return append([]int32{sep}, i...)
}

func DeleteInt32(slice []int32, sep int32) []int32 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutInt32(slice []int32, i, j int32) []int32 {
	return append(slice[:i], slice[j:]...)
}

func InsertInt32(slice []int32, element, position int32) []int32 {
	return append(slice[:position], append([]int32{element}, slice[position:]...)...)
}

func InsertVectorInt32(origin, insert []int32, position int32) []int32 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopInt32(slice []int32) (int32, []int32) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedInt32(slice []int32) []int32 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendInt32(A, B []int32) []int32 {
	return append(A, B...)
}
