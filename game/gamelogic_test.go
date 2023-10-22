package game

import (
	"fmt"

	"github.com/rhcourses-go/tictactoe/board"
)

// ExampleSwitchPlayer testet SwitchPlayer.
// Es wird erwartet, dass SwitchPlayer den Spieler wechselt.
// D.h. wenn der aktuelle Spieler X ist, wird O erwartet und umgekehrt.
func ExampleSwitchPlayer() {
	fmt.Println(SwitchPlayer("X"))
	fmt.Println(SwitchPlayer("O"))
	// Output:
	// O
	// X
}

// ExamplePlayerWins testet PlayerWins.
// Es werden verschiedene Spielsituationen getestet und erwartet,
// dass PlayerWins true liefert, wenn der Spieler gewonnen hat.
//
// Anmerkung: Die Test-Spielfelder in dieser Funktion müssen nicht
// unbedingt gültige Spielsituationen sein. Es geht hier nur darum,
// zu testen, ob PlayerWins die richtigen Antworten liefert.
func ExamplePlayerWins() {
	row_x := board.NewFromStrings(
		"XXX",
		"   ",
		"   ",
	)

	row_o := board.NewFromStrings(
		"   ",
		"OOO",
		"   ",
	)

	col_x := board.NewFromStrings(
		"X  ",
		"X  ",
		"X  ",
	)

	col_o := board.NewFromStrings(
		"  O",
		"  O",
		"  O",
	)

	diag_right_x := board.NewFromStrings(
		"X  ",
		" X ",
		"  X",
	)

	diag_left_o := board.NewFromStrings(
		"  O",
		" O ",
		"O  ",
	)

	draw := board.NewFromStrings(
		"XOX",
		"OXO",
		"OXO",
	)

	still_running := board.NewFromStrings(
		"XOX",
		"OX ",
		"OXO",
	)

	fmt.Println(PlayerWins("X", row_x)) // true
	fmt.Println(PlayerWins("O", row_x)) // false
	fmt.Println(PlayerWins("X", row_o)) // false
	fmt.Println(PlayerWins("O", row_o)) // true

	fmt.Println()

	fmt.Println(PlayerWins("X", col_x)) // true
	fmt.Println(PlayerWins("O", col_x)) // false
	fmt.Println(PlayerWins("X", col_o)) // false
	fmt.Println(PlayerWins("O", col_o)) // true

	fmt.Println()

	fmt.Println(PlayerWins("X", diag_right_x)) // true
	fmt.Println(PlayerWins("O", diag_right_x)) // false
	fmt.Println(PlayerWins("X", diag_left_o))  // false
	fmt.Println(PlayerWins("O", diag_left_o))  // true

	fmt.Println()

	fmt.Println(PlayerWins("X", draw))          // false
	fmt.Println(PlayerWins("O", draw))          // false
	fmt.Println(PlayerWins("X", still_running)) // false
	fmt.Println(PlayerWins("O", still_running)) // false

	// Output:
	// true
	// false
	// false
	// true
	//
	// true
	// false
	// false
	// true
	//
	// true
	// false
	// false
	// true
	//
	// false
	// false
	// false
	// false
}

// ExampleMoveAllowed testet MoveAllowed.
// Es werden verschiedene Spielsituationen getestet und erwartet,
// dass MoveAllowed true liefert, wenn der Zug erlaubt ist.
//
// Anmerkung: Die Test-Spielfelder in dieser Funktion müssen nicht
// unbedingt gültige Spielsituationen sein. Es geht hier nur darum,
// zu testen, ob MoveAllowed die richtigen Antworten liefert.
func ExampleMoveAllowed() {
	full := board.NewFromStrings(
		"XOX",
		"OXO",
		"XOX",
	)

	empty := board.NewFromStrings(
		"   ",
		"   ",
		"   ",
	)

	semi_full := board.NewFromStrings(
		"XOX",
		"   ",
		"OO ",
	)

	fmt.Println(MoveAllowed(full, 0))  // false
	fmt.Println(MoveAllowed(full, 1))  // false
	fmt.Println(MoveAllowed(full, 5))  // false
	fmt.Println(MoveAllowed(full, 9))  // false
	fmt.Println(MoveAllowed(full, 10)) // false

	fmt.Println()

	fmt.Println(MoveAllowed(empty, 0))  // false
	fmt.Println(MoveAllowed(empty, 1))  // true
	fmt.Println(MoveAllowed(empty, 5))  // true
	fmt.Println(MoveAllowed(empty, 9))  // true
	fmt.Println(MoveAllowed(empty, 10)) // false

	fmt.Println()

	fmt.Println(MoveAllowed(semi_full, 0))  // false
	fmt.Println(MoveAllowed(semi_full, 1))  // false
	fmt.Println(MoveAllowed(semi_full, 5))  // true
	fmt.Println(MoveAllowed(semi_full, 9))  // true
	fmt.Println(MoveAllowed(semi_full, 10)) // false

	// Output:
	// false
	// false
	// false
	// false
	// false
	//
	// false
	// true
	// true
	// true
	// false
	//
	// false
	// false
	// true
	// true
	// false
}
