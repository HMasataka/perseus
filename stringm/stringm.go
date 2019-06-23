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

func Shift(sep string, s []string) []string {
	return append([]string{sep}, s...)
}
