package intm

import ()

func In(i int, list []int) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func Index(s int, list []int) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func Shift(slice []int) (int, []int) {
	x, s := slice[0], slice[1:]
	return x, s
}

func Unshift(sep int, i []int) []int {
	return append([]int{sep}, i...)
}

func Delete(slice []int, sep int) []int {
	return append(slice[:sep], slice[sep+1:]...)
}

func Cut(slice []int, i, j int) []int {
	return append(slice[:i], slice[j:]...)
}

func Insert(slice []int, element, position int) []int {
	return append(slice[:position], append([]int{element}, slice[position:]...)...)
}

func Pop(slice []int) (int, []int) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func Reversed(slice []int) []int {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func Extend(A, B []int) []int {
	return append(A, B...)
}
