package boardrow

import "fmt"

func ExampleNew() {
	row := New(3, "A")
	for _, v := range row {
		fmt.Println(v)
	}
	// Output:
	// A
	// A
	// A
}

func ExampleNewEmpty() {
	row := NewEmpty(3)
	for _, v := range row {
		fmt.Println(v)
	}
	// Output:
	//
	//
	//
}

func ExampleNewFromString() {
	row := NewFromString("ABC")
	for _, v := range row {
		fmt.Println(v)
	}
	// Output:
	// A
	// B
	// C
}
