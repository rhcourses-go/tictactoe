package botgame

import (
	"github.com/rhcourses-go/tictactoe/board"
	"github.com/rhcourses-go/tictactoe/game"
)

// CopyBoard erwartet ein Board und liefert eine Kopie des Boards zurück.
func CopyBoard(b board.Board) board.Board {
	newBoard := board.New(len(b), len(b[0]), " ")
	/* Hinweis:
	   Verwenden Sie eine Schleife, um über alle Zeilen
	   des Spielfelds zu iterieren. In der Schleife können Sie
	   dann mit der Funktion copy die Zeilen kopieren.

	   Alternativ können Sie auch in einer geschachtelten Schleife
	   über alle Zeilen und Spalten iterieren und die Werte
	   einzeln kopieren.
	*/
	// tag::solution[]
	for i := range b {
		copy(newBoard[i], b[i])
	}
	// end::solution[]
	return newBoard
}

// WinningMove erwartet einen Spieler-String und ein Board und liefert
// die Zahl für einen Zug zurück, der zum Gewinn führt.
// Ist kein solcher Zug möglich, wird -1 zurückgegeben.
//
// Die Zahl liegt zwischen 1 und 9 und gibt das Feld an,
// auf dem der Zug gemacht werden muss, um zu gewinnen.
// Die Zählung beginnt links oben bei 1 und geht nach rechts und dann nach unten.
// Falls mehrere Züge zum Gewinn führen, wird einer davon zurückgegeben.
func WinningMove(player string, b board.Board) int {
	/* Hinweis:
	   Verwenden Sie eine Kopie des Spielfelds und einen Feld-Zähler,
	   um alle möglichen Züge zu testen. Laufen Sie dazu in einer geschachtelten
	   Schleife über alle Zeilen und Spalten des Spielfelds und ...

	   - prüfen Sie, ob der aktuelle Spieler auf das Feld setzen kann.
	   - setzen Sie den Spieler auf das Feld und prüfen Sie,
	     ob der Spieler gewonnen hat.
	   - liefern Sie die Nummer des Felds zurück, falls er gewonnen hat.
	   - setzen Sie das Feld wieder zurück, falls er nicht gewonnen hat.
	   - zählen Sie die Nummer des Felds hoch.

	   Um zu prüfen, ob ein Zug erlaubt ist oder ob ein Spieler gewonnen hat,
	   sollten Sie die Funktionen game.MoveAllowed und game.PlayerWins verwenden.
	*/
	// tag::solution[]
	boardcopy := CopyBoard(b)
	field := 1
	for row := range boardcopy {
		for col := range boardcopy[row] {
			if game.MoveAllowed(boardcopy, field) {
				boardcopy[row][col] = player
				if game.PlayerWins(player, boardcopy) {
					return field
				}
				boardcopy[row][col] = " "
			}
			field++
		}
	}
	// end::solution[]
	return -1
}

// BotMove erwartet einen Spieler-String und ein Board und liefert
// die Zahl für einen Zug zurück, den der Bot machen soll.
// Ist kein solcher Zug möglich, wird -1 zurückgegeben.
//
// Die Zahl liegt zwischen 1 und 9 und gibt das Feld an,
// auf dem der Zug gemacht werden soll.
//
// Der Bot versucht zu gewinnen, indem er die Funktion WinningMove verwendet.
// Falls kein Gewinn möglich ist, versucht er, den Gegner am Gewinnen zu hindern.
// Weitere Intelligenz ist nicht gefordert. Der Bot muss also nicht perfekt spielen.
func BotMove(player string, b board.Board) int {
	/* Hinweis:
	   Verwenden Sie die Funktion WinningMove, um zu prüfen,
	   ob der Bot gewinnen kann. Falls ja, geben Sie den Zug zurück.

	   Falls nicht, verwenden Sie die Funktion SwitchPlayer,
	   um den Gegner zu ermitteln. Verwenden Sie dann erneut
	   die Funktion WinningMove, um zu prüfen, ob der Gegner gewinnen kann.
	   Falls ja, besetzen Sie dieses Feld.

	   Falls beides nicht der Fall ist, liefern Sie das erste freie Feld zurück.
	   Falls es kein freies Feld gibt, liefern Sie -1 zurück.
	*/
	// tag::solution[]
	move := WinningMove(player, b)
	if move != -1 {
		return move
	}
	opponent := game.SwitchPlayer(player)
	move = WinningMove(opponent, b)
	if move != -1 {
		return move
	}
	for i := 1; i <= 9; i++ {
		if game.MoveAllowed(b, i) {
			return i
		}
	}
	// end::solution[]
	return -1
}
