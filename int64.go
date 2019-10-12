package perseus

func InInt64(i int64, list []int64) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexInt64(s int64, list []int64) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftInt64(slice []int64) (int64, []int64) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftInt64(sep int64, i []int64) []int64 {
	return append([]int64{sep}, i...)
}

func DeleteInt64(slice []int64, sep int64) []int64 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutInt64(slice []int64, i, j int64) []int64 {
	return append(slice[:i], slice[j:]...)
}

func InsertInt64(slice []int64, element, position int64) []int64 {
	return append(slice[:position], append([]int64{element}, slice[position:]...)...)
}

func InsertVectorInt64(origin, insert []int64, position int64) []int64 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopInt64(slice []int64) (int64, []int64) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedInt64(slice []int64) []int64 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendInt64(A, B []int64) []int64 {
	return append(A, B...)
}
