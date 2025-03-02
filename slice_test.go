package perseus_test

import (
	"fmt"
	"testing"

	"github.com/HMasataka/perseus"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	x, s := perseus.Shift(v)
	assert.Equal(t, []int{2, 3, 4, 5}, s)

	fmt.Printf("Shift : %d, %v\n", x, s)
	// Output:
	// Shift : 1, [2 3 4 5]
}

func TestPrepend(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	u := perseus.Prepend(1, v)
	assert.Equal(t, []int{1, 1, 2, 3, 4, 5}, u)

	fmt.Printf("Prepend : %v\n", u)
	// Output:
	// Prepend : [1 1 2 3 4 5]
}

func TestRemove(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	d := perseus.Remove(v, 3)
	assert.Equal(t, []int{1, 2, 3, 5}, d)

	fmt.Printf("Remove : %v\n", d)
	// Output:
	// Remove : [1 2 3 5]
}

func TestCut(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	c := perseus.Cut(v, 2, 3)
	assert.Equal(t, []int{1, 2, 4, 5}, c)

	fmt.Printf("Cut : %v\n", c)
	// Output:
	// Cut : [1 2 4 5]

}

func TestInsert(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	insert := perseus.Insert(v, 10, 3)
	assert.Equal(t, []int{1, 2, 3, 10, 4, 5}, insert)

	fmt.Printf("Insert : %v\n", insert)
	// Output:
	// Insert : [1 2 3 10 4 5]
}

func TestInsertVector(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	insertV := perseus.InsertVector(v, []int{15, 14, 13, 12}, 3)
	assert.Equal(t, []int{1, 2, 3, 15, 14, 13, 12, 4, 5}, insertV)

	fmt.Printf("InsertVector : %v\n", insertV)
	// Output:
	// InsertVector : [1 2 3 15 14 13 12 4 5]
}

func TestPop(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	p, vv := perseus.Pop(v)
	assert.Equal(t, []int{1, 2, 3, 4}, vv)

	fmt.Printf("Pop : %d, %v\n", p, vv)
	// Output:
	// Pop : 5, [1 2 3 4]
}

func TestExtend(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	e := perseus.Extend(v, []int{11, 12, 13, 14, 15})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 11, 12, 13, 14, 15}, e)

	fmt.Printf("Extend : %v\n", e)
	// Output:
	// Extend : [1 2 3 4 5 11 12 13 14 15]
}

func TestSum(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	e := perseus.Sum(v)
	assert.Equal(t, 15, e)

	fmt.Printf("Sum : %v\n", e)
	// Output:
	// Extend :[15
}
