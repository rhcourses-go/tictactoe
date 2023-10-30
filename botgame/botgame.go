package botgame

import (
	"fmt"

	"github.com/rhcourses-go/tictactoe/board"
	"github.com/rhcourses-go/tictactoe/game"
)

func Run() {

	// Spielfeld initialisieren.
	board := board.NewFromStrings(
		"123",
		"456",
		"789",
	)

	// Spieler initialisieren.
	currentPlayer := "O"

	// Zug-Zähler initialisieren.
	moveCount := 0

	// Hauptschleife.
	for moveCount < 9 {
		// Prüfen, ob der aktuelle Spieler gewonnen hat.
		if game.PlayerWins(currentPlayer, board) {
			fmt.Println("Spieler", currentPlayer, "gewinnt!")
			return
		}

		// Spieler wechseln und Zug machen.
		currentPlayer = game.SwitchPlayer(currentPlayer)

		// Zug machen und Zug-Zähler erhöhen.
		/* Hinweis:
		   Entscheiden Sie je nach aktuellem Spieler,
		   ob der Zug vom Menschen oder vom Bot gemacht wird.
		   Verwenden Sie dazu die Funktionen game.MakeMove und BotMove.
		   Falls der Bot am Zug ist, führen Sie seinen Zug aus.
		*/
		// tag::solution[]
		if currentPlayer == "X" {
			game.MakeMove(currentPlayer, board)
		} else {
			bm := BotMove(currentPlayer, board)
			board[(bm-1)/3][(bm-1)%3] = currentPlayer
		}
		// end::solution[]
		moveCount++
	}

	// Prüfen, ob das Spiel unentschieden ist und das Ergebnis ausgeben.
	if game.PlayerWins(currentPlayer, board) {
		fmt.Println("Spieler", currentPlayer, "gewinnt!")
	} else {
		fmt.Println("Unentschieden!")
	}
}
