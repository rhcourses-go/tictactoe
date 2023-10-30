package main

import (
	"fmt"

	"github.com/rhcourses-go/tictactoe/botgame"
)

func main() {
	fmt.Println("Willkommen bei Tic Tac Bot!")
	fmt.Println("Das Spiel beginnt...")
	fmt.Println()
	botgame.Run()
}
