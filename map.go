package perseus

import "iter"

type KeyValueIndex[K comparable, V any] struct {
	Index int
	Key   K
	Value V
}

func WithIndex[K comparable, V any](m map[K]V) iter.Seq[KeyValueIndex[K, V]] {
	return func(yield func(KeyValueIndex[K, V]) bool) {
		idx := 0
		for key, value := range m {
			yield(KeyValueIndex[K, V]{
				Index: idx,
				Key:   key,
				Value: value,
			})
			idx++
		}
	}
}
