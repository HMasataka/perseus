package perseus

func InFloat64(i float64, list []float64) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexFloat64(s float64, list []float64) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

func ShiftFloat64(slice []float64) (float64, []float64) {
	x, s := slice[0], slice[1:]
	return x, s
}

func UnshiftFloat64(sep float64, i []float64) []float64 {
	return append([]float64{sep}, i...)
}

func DeleteFloat64(slice []float64, sep int) []float64 {
	return append(slice[:sep], slice[sep+1:]...)
}

func CutFloat64(slice []float64, i, j int) []float64 {
	return append(slice[:i], slice[j:]...)
}

func InsertFloat64(slice []float64, element float64, position int) []float64 {
	return append(slice[:position], append([]float64{element}, slice[position:]...)...)
}

func InsertVectorFloat64(origin, insert []float64, position int) []float64 {
	return append(origin[:position], append(insert, origin[position:]...)...)
}

func PopFloat64(slice []float64) (float64, []float64) {
	x, s := slice[len(slice)-1], slice[:len(slice)-1]
	return x, s
}

func ReversedFloat64(slice []float64) []float64 {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
	return slice
}

func ExtendFloat64(A, B []float64) []float64 {
	return append(A, B...)
}
