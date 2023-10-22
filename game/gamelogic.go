package game

import (
	"github.com/rhcourses-go/tictactoe/board"
)

// SwitchPlayer erwartet einen Spieler-String und liefert den anderen Spieler zurück.
func SwitchPlayer(currentPlayer string) string {
	/* Hinweis:
	   Prüfen Sie, ob der aktuelle Spieler "X" ist.
	   Wenn ja, geben Sie "O" zurück, sonst "X".
	*/
	// tag::solution[]
	if currentPlayer == "X" {
		return "O"
	}
	// end::solution[]
	return "X"
}

// PlayerWins erwartet einen Spieler-String und ein Board und liefert true zurück,
// wenn der Spieler gewonnen hat.
func PlayerWins(player string, b board.Board) bool {
	/* Hinweis:
	   Prüfen Sie, ob der Spieler in einer Reihe, Spalte
	   oder Diagonale alle Felder belegt hat.
	   Dafür können Sie in einer Schleife über 0...2 iterieren
	   und die Funktionen RowContainsOnly und ColContainsOnly verwenden.
	   Die Diagonalen sollten Sie separat prüfen.
	*/
	// tag::solution[]
	for i := range b {
		if board.RowContainsOnly(b, i, player) || board.ColumnContainsOnly(b, i, player) {
			return true
		}
	}
	if board.DiagRightContainsOnly(b, player) || board.DiagLeftContainsOnly(b, player) {
		return true
	}
	// end::solution[]
	return false
}

// MoveAllowed erwartet ein Board und eine Zahl und liefert true zurück,
// wenn das Feld noch nicht belegt ist.
//
// Die Zahl muss zwischen 1 und 9 liegen und gibt das Feld an,
// das geprüft werden soll. Die Zählung beginnt links oben bei 1
// und geht nach rechts und dann nach unten.
//
// Liegt die Zahl außerhalb des Bereichs 1-9, wird false zurückgegeben.
func MoveAllowed(b board.Board, move int) bool {
	/* Hinweis:
	   Prüfen Sie, ob die Zahl zwischen 1 und 9 liegt.
	   Wenn nicht, geben Sie false zurück.
	   Ansonsten überlegen Sie sich, wie Sie aus der Zahl
	   die Reihe und Spalte berechnen können.
	   Auf das Spielfeld an Stelle [row][col] können Sie
	   dann mittels `b[row][col]` zugreifen.

	   Alternativ können Sie auch in einer Doppel-Schleife
	   über das Spielfeld iterieren und dabei mitzählen,
	   um das richtige Feld zu finden.
	*/
	// tag::solution[]
	if move < 1 || move > 9 {
		return false
	}
	row := (move - 1) / 3
	col := (move - 1) % 3
	return b[row][col] != "X" && b[row][col] != "O"
	// end::solution[]
	// iftask: return false
}
