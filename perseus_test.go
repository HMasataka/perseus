package perseus_test

import (
	"fmt"

	"github.com/sylba2050/perseus"
)

func ExampleInInt() {
	v := []int{1, 2, 3, 4, 5}

	b := perseus.InInt(1, v)
	fmt.Printf("In : %t\n", b)
	// Output:
	// In : true
}

func ExampleIndexInt() {
	v := []int{1, 2, 3, 4, 5}

	i := perseus.IndexInt(3, v)
	fmt.Printf("Index : %d\n", i)
	// Output:
	// Index : 2
}

func ExampleShiftInt() {
	v := []int{1, 2, 3, 4, 5}

	x, s := perseus.ShiftInt(v)
	fmt.Printf("Shift : %d, %v\n", x, s)
	// Output:
	// Shift : 1, [2 3 4 5]
}

func ExampleUnshiftInt() {
	v := []int{1, 2, 3, 4, 5}

	u := perseus.UnshiftInt(1, v)
	fmt.Printf("Unshift : %v\n", u)
	// Output:
	// Unshift : [1 1 2 3 4 5]
}

func ExampleDeleteInt() {
	v := []int{1, 2, 3, 4, 5}

	d := perseus.DeleteInt(v, 3)
	fmt.Printf("Delete : %v\n", d)
	// Output:
	// Delete : [1 2 3 5]
}
func ExampleCutInt() {
	v := []int{1, 2, 3, 4, 5}

	c := perseus.CutInt(v, 2, 3)
	fmt.Printf("Cut : %v\n", c)
	// Output:
	// Cut : [1 2 4 5]

}
func ExampleInsertInt() {
	v := []int{1, 2, 3, 4, 5}

	insert := perseus.InsertInt(v, 10, 3)
	fmt.Printf("Insert : %v\n", insert)
	// Output:
	// Insert : [1 2 3 10 4 5]
}

func ExampleInsertVectorInt() {
	v := []int{1, 2, 3, 4, 5}

	insertV := perseus.InsertVectorInt(v, []int{15, 14, 13, 12}, 3)
	fmt.Printf("InsertV : %v\n", insertV)
	// Output:
	// InsertV : [1 2 3 15 14 13 12 4 5]
}

func ExamplePopInt() {
	v := []int{1, 2, 3, 4, 5}

	p, vv := perseus.PopInt(v)
	fmt.Printf("Pop : %d, %v\n", p, vv)
	// Output:
	// Pop : 5, [1 2 3 4]
}

func ExampleReversedInt() {
	v := []int{1, 2, 3, 4, 5}

	r := perseus.ReversedInt(v)
	fmt.Printf("Reversed : %v\n", r)
	// Output:
	// Reversed : [5 4 3 2 1]
}

func ExampleExtendInt() {
	v := []int{1, 2, 3, 4, 5}
	e := perseus.ExtendInt(v, []int{11, 12, 13, 14, 15})
	fmt.Printf("Extend : %v\n", e)
	// Output:
	// Extend : [1 2 3 4 5 11 12 13 14 15]
}
