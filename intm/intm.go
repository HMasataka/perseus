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

func Shift(sep int, i []int) []int {
	return append([]int{sep}, i...)
}
