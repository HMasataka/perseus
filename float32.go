package perseus

// InFloat32 returns whether i is in slice.
func InFloat32(i float32, slice []float32) bool {
	for _, b := range slice {
		if b == i {
			return true
		}
	}
	return false
}

// IndexFloat32 returns the position of s in slice. If s is not found, return -1.
func IndexFloat32(s float32, slice []float32) int {
	for i, b := range slice {
		if b == s {
			return i
		}
	}
	return -1
}

// ShiftFloat32 returns the first element of slice and other element's slice.
func ShiftFloat32(slice []float32) (float32, []float32) {
	return slice[0], slice[1:]
}

// UnshiftFloat32 add an element to the beginning of a slice.
func UnshiftFloat32(sep float32, i []float32) []float32 {
	return append([]float32{sep}, i...)
}

// DeleteFloat32 Delete specified element from slice
func DeleteFloat32(slice []float32, sep int) []float32 {
	return append(slice[:sep], slice[sep+1:]...)
}

// CutFloat32 Delete from i to j from the slice
func CutFloat32(slice []float32, i, j int) []float32 {
	return append(slice[:i], slice[j:]...)
}

// InsertFloat32 Insert element to specified position
func InsertFloat32(slice []float32, element float32, position int) []float32 {
	return append(slice[:position], append([]float32{element}, slice[position:]...)...)
}

// InsertVectorFloat32 Insert slice to specified position
func InsertVectorFloat32(origin, insert []float32, position int) []float32 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

// PopFloat32 returns the last element of slice and other element's slice.
func PopFloat32(slice []float32) (float32, []float32) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// ReversedFloat32 returns reversed slice
func ReversedFloat32(slice []float32) []float32 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

// ExtendFloat32 connect slices together
func ExtendFloat32(A, B []float32) []float32 {
	return append(A, B...)
}
