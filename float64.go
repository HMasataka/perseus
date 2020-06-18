package perseus

// InFloat64 returns whether i is in list
func InFloat64(i float64, list []float64) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

// IndexFloat64 returns the position of s in list. If s is not found, return -1.
func IndexFloat64(s float64, list []float64) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftFloat64 returns the first element of slice and other element's slice.
func ShiftFloat64(slice []float64) (float64, []float64) {
	return slice[0], slice[1:]
}

// UnshiftFloat64 add an element to the beginning of a slice.
func UnshiftFloat64(sep float64, i []float64) []float64 {
	return append([]float64{sep}, i...)
}

// DeleteFloat64 Delete specified element from slice
func DeleteFloat64(slice []float64, sep int) []float64 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutFloat64 Delete from i to j from the slice
func CutFloat64(slice []float64, i, j int) []float64 {
	return append(slice[:i], slice[j:]...)
}

// InsertFloat64 Insert element to specified position
func InsertFloat64(slice []float64, element float64, position int) []float64 {
	return append(slice[:position], append([]float64{element}, slice[position:]...)...)
}

// InsertVectorFloat64 Insert slice to specified position
func InsertVectorFloat64(origin, insert []float64, position int) []float64 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopFloat64 returns the last element of slice and other element's slice.
func PopFloat64(slice []float64) (float64, []float64) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// ReversedFloat64 returns reversed slice
func ReversedFloat64(slice []float64) []float64 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendFloat64 connect slices together
func ExtendFloat64(A, B []float64) []float64 {
	return append(A, B...)
}

// SumFloat64 calculate summaries of arguments
func SumFloat64(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}
