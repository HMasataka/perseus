package stringm

import ()

func In(s string, list []string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

func Index(s string, list []string) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func Shift(slice []string) (string, []string) {
	x, s := slice[0], slice[1:]
	return x, s
}

func Unshift(sep string, s []string) []string {
	return append([]string{sep}, s...)
}

func Delete(slice []string, sep int) []string {
	return append(slice[:sep], slice[sep+1:]...)
}

func Cut(slice []string, i, j int) []string {
	return append(slice[:i], slice[j:]...)
}

func Insert(slice []string, element string, position int) []string {
	return append(slice[:position], append([]string{element}, slice[position:]...)...)
}

func Pop(slice []string) (string, []string) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func Reversed(slice []string) []string {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func Extend(A, B []string) []string {
	return append(A, B...)
}
