package board

import "fmt"

func ExampleBoardRowContainsOnly() {
	board := NewEmpty(3, 2)
	board[0][0] = "A"
	board[0][1] = "A"
	board[0][2] = "A"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"

	fmt.Println(RowContainsOnly(board, 0, "A"))
	fmt.Println(RowContainsOnly(board, 0, "B"))
	fmt.Println(RowContainsOnly(board, 1, "D"))
	fmt.Println(RowContainsOnly(board, 1, "E"))

	// Output:
	// true
	// false
	// false
	// false
}

func ExampleBoardColumnContainsOnly() {
	board := NewEmpty(3, 2)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "A"
	board[1][1] = "E"
	board[1][2] = "C"

	fmt.Println(ColumnContainsOnly(board, 0, "A"))
	fmt.Println(ColumnContainsOnly(board, 0, "B"))
	fmt.Println(ColumnContainsOnly(board, 1, "B"))
	fmt.Println(ColumnContainsOnly(board, 1, "E"))
	fmt.Println(ColumnContainsOnly(board, 2, "C"))
	fmt.Println(ColumnContainsOnly(board, 2, "F"))

	// Output:
	// true
	// false
	// false
	// false
	// true
	// false
}

func ExampleBoardDiagRightContainsOnly() {
	board := NewEmpty(3, 3)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "A"
	board[1][2] = "F"
	board[2][0] = "G"
	board[2][1] = "H"
	board[2][2] = "A"

	fmt.Println(DiagRightContainsOnly(board, "A"))
	fmt.Println(DiagRightContainsOnly(board, "B"))

	// Output:
	// true
	// false
}

func ExampleBoardDiagLeftContainsOnly() {
	board := NewEmpty(3, 3)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "C"
	board[1][2] = "F"
	board[2][0] = "C"
	board[2][1] = "H"
	board[2][2] = "I"

	fmt.Println(DiagLeftContainsOnly(board, "C"))
	fmt.Println(DiagLeftContainsOnly(board, "B"))

	// Output:
	// true
	// false
}

func ExampleBoardContainsOnly() {
	board1 := NewEmpty(3, 2)
	fmt.Println(ContainsOnly(board1, " "))
	fmt.Println(ContainsOnly(board1, "A"))
	board1[0][0] = "A"
	fmt.Println(ContainsOnly(board1, " "))

	board2 := New(3, 2, "A")
	fmt.Println(ContainsOnly(board2, "A"))
	fmt.Println(ContainsOnly(board2, "B"))
	board2[0][0] = "B"
	fmt.Println(ContainsOnly(board2, "A"))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// false
}

func ExampleBoardDoesNotContain() {
	board1 := NewEmpty(3, 2)
	fmt.Println(ContainsNone(board1, " "))
	fmt.Println(ContainsNone(board1, "A"))
	board1[0][0] = "A"
	fmt.Println(ContainsNone(board1, " "))
	fmt.Println(ContainsNone(board1, "A"))
	board1[0][1] = "B"
	board1[0][2] = "C"
	board1[1][0] = "D"
	board1[1][1] = "E"
	board1[1][2] = "F"
	fmt.Println(ContainsNone(board1, " "))
	fmt.Println(ContainsNone(board1, "A"))
	fmt.Println(ContainsNone(board1, "B"))
	fmt.Println(ContainsNone(board1, "K"))

	// Output:
	// false
	// true
	// false
	// false
	// true
	// false
	// false
	// true
}
