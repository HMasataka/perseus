package perseus

func InFloat32(i float32, list []float32) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexFloat32(s float32, list []float32) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftFloat32(slice []float32) (float32, []float32) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftFloat32(sep float32, i []float32) []float32 {
	return append([]float32{sep}, i...)
}

func DeleteFloat32(slice []float32, sep int) []float32 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutFloat32(slice []float32, i, j int) []float32 {
	return append(slice[:i], slice[j:]...)
}

func InsertFloat32(slice []float32, element float32, position int) []float32 {
	return append(slice[:position], append([]float32{element}, slice[position:]...)...)
}

func InsertVectorFloat32(origin, insert []float32, position int) []float32 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopFloat32(slice []float32) (float32, []float32) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedFloat32(slice []float32) []float32 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendFloat32(A, B []float32) []float32 {
	return append(A, B...)
}
