package boardrow

func ExamplePrint() {
	row := New(3, "A")
	Print(row)

	// Output:
	// | A | A | A |
}
