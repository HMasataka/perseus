package perseus

func InInt(i int, list []int) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexInt(s int, list []int) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftInt(slice []int) (int, []int) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftInt(sep int, i []int) []int {
	return append([]int{sep}, i...)
}

func DeleteInt(slice []int, sep int) []int {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutInt(slice []int, i, j int) []int {
	return append(slice[:i], slice[j:]...)
}

func InsertInt(slice []int, element, position int) []int {
	return append(slice[:position], append([]int{element}, slice[position:]...)...)
}

func InsertVectorInt(origin, insert []int, position int) []int {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopInt(slice []int) (int, []int) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedInt(slice []int) []int {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendInt(A, B []int) []int {
	return append(A, B...)
}
