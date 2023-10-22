package main

import (
	"fmt"

	"github.com/rhcourses-go/tictactoe/board"
	"github.com/rhcourses-go/tictactoe/game"
)

func main() {
	// Je nach Bedarf ein- oder auskommentieren.

	TestAskForMove()
	// TestMakeMove()
}

func TestAskForMove() {
	result := game.AskForMove("X")
	fmt.Printf("AskForMove(\"X\") == %d\n", result)
}

func TestMakeMove() {
	b := board.NewFromStrings(
		"OX ",
		" O ",
		"  X",
	)

	game.MakeMove("X", b)
	fmt.Printf("MakeMove(\"X\", b) == \n%s\n", b)
}
