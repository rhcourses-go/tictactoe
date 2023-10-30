package botgame

import (
	"fmt"

	"github.com/rhcourses-go/tictactoe/board"
)

// ExampleCopyBoard zeigt die Verwendung der Funktion CopyBoard.
// Es wird ein Spielfeld erzeugt und mit CopyBoard kopiert.
// Anschließend wird die Kopie verändert und das Original ausgegeben.
// Die Ausgabe zeigt, dass das Original unverändert bleibt.
func ExampleCopyBoard() {
	b := board.NewFromStrings(
		"ABC",
		"DEF",
		"GHI",
	)
	bcopy := CopyBoard(b)
	bcopy[0][0] = "1"
	bcopy[0][1] = "2"
	bcopy[0][2] = "3"
	bcopy[1][0] = "1"
	bcopy[1][1] = "2"
	bcopy[1][2] = "3"
	bcopy[2][0] = "1"
	bcopy[2][1] = "2"
	bcopy[2][2] = "3"

	for _, row := range b {
		fmt.Println([]string(row))
	}

	// Output:
	// [A B C]
	// [D E F]
	// [G H I]
}

// ExampleWinningMove zeigt die Verwendung der Funktion WinningMove.
// Es werden verschiedene Spielfelder erzeugt und die Funktion WinningMove
// mit diesen Spielfeldern aufgerufen.
// Die Ausgabe zeigt, dass die Funktion WinningMove die richtigen Züge liefert.
func ExampleWinningMove() {
	// Spielfeld erzeugen.
	both_win_1 := board.NewFromStrings(
		"X O",
		"XO ",
		"   ",
	)
	fmt.Println(WinningMove("X", both_win_1))
	fmt.Println(WinningMove("O", both_win_1))
	fmt.Println()

	both_win_2 := board.NewFromStrings(
		"XX ",
		"   ",
		"O O",
	)
	fmt.Println(WinningMove("X", both_win_2))
	fmt.Println(WinningMove("O", both_win_2))
	fmt.Println()

	x_wins := board.NewFromStrings(
		"XO ",
		"X O",
		"   ",
	)
	fmt.Println(WinningMove("X", x_wins))
	fmt.Println(WinningMove("O", x_wins))
	fmt.Println()

	o_wins := board.NewFromStrings(
		"OX ",
		"X  ",
		"  O",
	)
	fmt.Println(WinningMove("X", o_wins))
	fmt.Println(WinningMove("O", o_wins))
	fmt.Println()

	none_wins := board.NewFromStrings(
		"XO ",
		"  O",
		" X ",
	)
	fmt.Println(WinningMove("X", none_wins))
	fmt.Println(WinningMove("O", none_wins))

	// Output:
	// 7
	// 7
	//
	// 3
	// 8
	//
	// 7
	// -1
	//
	// -1
	// 5
	//
	// -1
	// -1
}

// ExampleBotMove_x_wins verwendet ein Spielfeld, in dem Spieler X gewinnen kann.
// Das Beispiel zeigt, dass der Bot für beide Spieler den richtigen Zug (7) liefert.
// Als Spieler X gewinnt er so, als Spieler O verhindert er den Sieg von X.
func ExampleBotMove_x_wins() {
	b := board.NewFromStrings(
		"XO ",
		"X O",
		"   ",
	)
	fmt.Println(BotMove("X", b))
	fmt.Println(BotMove("O", b))

	// Output:
	// 7
	// 7
}

// ExampleBotMove_o_wins verwendet ein Spielfeld, in dem Spieler O gewinnen kann.
// Das Beispiel zeigt, dass der Bot für beide Spieler den richtigen Zug (5) liefert.
// Als Spieler O gewinnt er so, als Spieler X verhindert er den Sieg von O.
func ExampleBotMove_o_wins() {
	b := board.NewFromStrings(
		"OX ",
		"X  ",
		"  O",
	)
	fmt.Println(BotMove("X", b))
	fmt.Println(BotMove("O", b))

	// Output:
	// 5
	// 5
}

// ExampleBotMove_none_wins verwendet ein Spielfeld, in dem keiner der Spieler gewinnen kann.
// Das Beispiel zeigt, dass der Bot für beide Spieler das erste freie Feld (3) liefert.
func ExampleBotMove_none_wins() {
	b := board.NewFromStrings(
		"XO ",
		"  O",
		" X ",
	)
	fmt.Println(BotMove("X", b))
	fmt.Println(BotMove("O", b))

	// Output:
	// 3
	// 3
}

// ExampleBotMove_full verwendet ein volles Spielfeld.
// Das Beispiel zeigt, dass der Bot für beide Spieler -1 liefert.
func ExampleBotMove_full() {
	b := board.NewFromStrings(
		"XOX",
		"OXO",
		"XOX",
	)
	fmt.Println(BotMove("X", b))
	fmt.Println(BotMove("O", b))

	// Output:
	// -1
	// -1
}
