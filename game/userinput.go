package game

/******************************************************************
 *                                                                *
 * Anmerkung zu den Funktionen in dieser Datei:                   *
 *                                                                *
 * Die Funktionen in dieser Datei beinhalten Benutzereingaben.    *
 * Dies kann nicht auf einfache Weise simuliert werden und daher  *
 * gibt es auch keine automatisierten Tests für diese Funktionen. *
 *                                                                *
 * Diese Funktionen sollten daher manuell getestet werden.        *
 * Dies kann z.B. in einem eigenen `main`-Package geschehen.      *
 * Zu diesem Zweck gibt es ein separates Package `testinput`.     *
 *                                                                *
 * Es kann auch sinnvoll sein, zuerst die Funktionen in           *
 * gamelogic_test.go` zu implementieren                           *
 *                                                                *
 ******************************************************************/

import (
	"fmt"

	"github.com/rhcourses-go/tictactoe/board"
)

// MakeMove erwartet einen Spieler-String und ein Board und lässt den Spieler
// einen Zug machen.
// Dazu zeigt die Funktion dem Spieler das Spielfeld und fragt den Spieler
// nach einer Zahl für das Feld. Anschließend setzt sie das entsprechende Feld
// auf den Spieler-String.
//
// Ist der Zug ungültig, wird der Spieler erneut nach einer Zahl gefragt.
func MakeMove(player string, b board.Board) {
	move := -1
	fmt.Println("Aktuelles Spielfeld:")
	board.PrintBoard(b)

	/* Hinweis:
	   Verwenden Sie eine Schleife, um den Spieler so lange nach einem Zug
	   zu fragen, bis er einen gültigen Zug macht.
	   D.h. die Schleife läuft so lange, bis die Funktion AskForMove
	   einen Zug liefert, für den MoveAllowed true liefert.
	*/
	// tag::AskForMove[]
	for move == -1 {
		move = AskForMove(player)
		if !MoveAllowed(b, move) {
			fmt.Println("Dieser Zug ist nicht erlaubt!")
			move = -1
		}
	}
	// end::AskForMove[]

	/* Hinweis:
	   Berechnen Sie aus der Zahl den Index für die Reihe und Spalte.
	   Setzen Sie dann das entsprechende Feld auf den Spieler-String.
	   Auf das Spielfeld an Stelle [row][col] können Sie
	   mittels `b[row][col]` zugreifen.
	*/
	// tag::doMove[]
	row := (move - 1) / 3
	col := (move - 1) % 3
	b[row][col] = player
	// end::doMove[]
}

// AskForMove fragt den Spieler nach einer Zahl für ein Feld
// auf dem Spielfeld und liefert diese zurück.
// Ist die Eingabe ungültig, wird -1 zurückgegeben.
func AskForMove(player string) int {
	var move string

	/* Hinweis:
	   Geben Sie aus, welcher Spieler am Zug ist und
	   fragen Sie nach einer Zahl für ein Feld.
	   Verwenden Sie fmt.Scan, um die Eingabe des Spielers einzulesen.
	   Die Eingabe wird in der Variable `move` gespeichert.

	   Prüfen Sie anschließend, ob die Eingabe gültig ist.
	   D.h. prüfen Sie, ob die Eingabe einer der Strings "1", "2", ... "9" ist.
	   Falls ja, geben Sie die Position der Eingabe in der Reihe
	   "1" ... "9" zurück. D.h. "1" wird zu 0, "2" zu 1 usw.
	*/
	// tag::solution[]
	fmt.Print("Spieler ", player, " ist am Zug. Bitte ein Feld eingeben: ")
	fmt.Scan(&move)

	allowedInputs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i, m := range allowedInputs {
		if m == move {
			return i + 1
		}
	}
	// end::solution[]
	return -1 // ungültige Eingabe
}
