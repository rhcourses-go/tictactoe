package boardrow

import "fmt"

func ExampleContainsOnly() {
	row1 := New(3, "A")
	fmt.Println(ContainsOnly(row1, "A"))
	fmt.Println(ContainsOnly(row1, "B"))

	row2 := NewEmpty(3)
	fmt.Println(ContainsOnly(row2, " "))
	row2[1] = "A"
	fmt.Println(ContainsOnly(row2, " "))

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleIsEmpty() {
	row1 := New(3, "A")
	fmt.Println(IsEmpty(row1))

	row2 := NewEmpty(3)
	fmt.Println(IsEmpty(row2))

	// Output:
	// false
	// true
}
