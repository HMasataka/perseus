package perseus

import "golang.org/x/exp/constraints"

// Shift returns the first element of slice and other element's slice.
func Shift[T any](slice []T) (T, []T) {
	return slice[0], slice[1:]
}

// Prepend add an element to the beginning of a slice.
func Prepend[T any](sep T, i []T) []T {
	return append([]T{sep}, i...)
}

// Remove Remove specified element from slice
func Remove[K any, V constraints.Integer](slice []K, sep V) []K {
	return append(slice[:sep], slice[sep+1:]...)
}

// Cut Delete from i to j from the slice
func Cut[K any, V constraints.Integer](slice []K, i, j V) []K {
	return append(slice[:i], slice[j:]...)
}

// Insert Insert element to specified position
func Insert[K any, V constraints.Integer](slice []K, element K, position V) []K {
	return append(slice[:position], append([]K{element}, slice[position:]...)...)
}

// InsertVector Insert slice to specified position
func InsertVector[K any, V constraints.Integer](origin, insert []K, position V) []K {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// Pop returns the last element of slice and other element's slice.
func Pop[T any](slice []T) (T, []T) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// Extend connect slices together
func Extend[T any](A, B []T) []T {
	return append(A, B...)
}

func Sum[T Number](values []T) T {
	var sum T

	for _, v := range values {
		sum += v
	}

	return sum
}

func Repeat[T any](count int, v ...T) []T {
	l := count * len(v)
	xs := make([]T, l)

	for i := 0; i < l; i += len(v) {
		copy(xs[i:], v)
	}

	return xs
}
