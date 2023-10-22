package board

import "fmt"

func ExampleNew() {
	board := New(3, 2, "A")
	for _, row := range board {
		fmt.Println(row)
	}

	// Output:
	// [A A A]
	// [A A A]
}

func ExampleNewEmpty() {
	board := NewEmpty(3, 2)
	for _, row := range board {
		fmt.Println(row)
	}

	// Output:
	// [     ]
	// [     ]
}

func ExampleNewFromStrings() {
	board := NewFromStrings(
		"ABC",
		"DEF",
	)
	for _, row := range board {
		fmt.Println(row)
	}

	// Output:
	// [A B C]
	// [D E F]
}

func ExampleRow() {
	board := NewEmpty(3, 2)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"

	fmt.Println(Row(board, 0))
	fmt.Println(Row(board, 1))

	// Output:
	// [A B C]
	// [D E F]
}

func ExampleColumn() {
	board := NewEmpty(3, 2)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"

	fmt.Println(Column(board, 0))
	fmt.Println(Column(board, 1))
	fmt.Println(Column(board, 2))

	// Output:
	// [A D]
	// [B E]
	// [C F]
}

func ExampleDiagRight() {
	board := NewEmpty(3, 3)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"
	board[2][0] = "G"
	board[2][1] = "H"
	board[2][2] = "I"

	fmt.Println(DiagRight(board))

	// Output:
	// [A E I]
}

func ExampleDiagLeft() {
	board := NewEmpty(3, 3)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"
	board[2][0] = "G"
	board[2][1] = "H"
	board[2][2] = "I"

	fmt.Println(DiagLeft(board))

	// Output:
	// [C E G]
}
