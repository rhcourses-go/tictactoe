package board

import "fmt"

func ExampleRowDivider() {
	fmt.Println(RowDivider(2))
	fmt.Println(RowDivider(3))
	fmt.Println(RowDivider(4))

	// Output:
	// +---+---+
	// +---+---+---+
	// +---+---+---+---+
}

func ExamplePrintBoard() {
	board := NewEmpty(3, 2)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"

	PrintBoard(board)

	// Output:
	// +---+---+---+
	// | A | B | C |
	// +---+---+---+
	// | D | E | F |
	// +---+---+---+
}
