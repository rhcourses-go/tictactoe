package main

import (
	"fmt"

	"github.com/rhcourses-go/tictactoe/game"
)

func main() {
	fmt.Println("Willkommen bei Tic Tac Toe!")
	fmt.Println("Das Spiel beginnt...")
	fmt.Println()
	game.Run()
}
