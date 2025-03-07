package perseus_test

import (
	"fmt"
	"testing"

	"github.com/HMasataka/perseus"
	"github.com/stretchr/testify/assert"
)

func TestWithIndex(t *testing.T) {
	v := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	w := perseus.WithIndex(v)

	i := 0
	for kv := range w {
		fmt.Printf("Index : %d, Key : %s, Value : %d\n", kv.Index, kv.Key, kv.Value)

		assert.Equal(t, kv.Index, i)
		assert.Equal(t, v[kv.Key], kv.Value)

		i++
	}
	// Output:
	// Index : 0, Key : four, Value : 4
	// Index : 1, Key : five, Value : 5
	// Index : 2, Key : one, Value : 1
	// Index : 3, Key : two, Value : 2
	// Index : 4, Key : three, Value : 3
}
