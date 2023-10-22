package board

import "github.com/rhcourses-go/tictactoe/boardrow"

// RowContainsOnly prüft, ob die angegebene Zeile nur aus dem übergebenen String besteht.
func RowContainsOnly(b Board, row int, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion Row.
	   Verwenden Sie die Funktion boardrow.ContainsOnly.
	*/
	// tag::solution[]
	return boardrow.ContainsOnly(Row(b, row), s)
	// end::solution[]
	// iftask: return false
}

// ColumnContainsOnly prüft, ob die angegebene Spalte nur aus dem übergebenen String besteht.
func ColumnContainsOnly(b Board, column int, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion Column.
	   Verwenden Sie die Funktion boardrow.ContainsOnly.
	*/
	// tag::solution[]
	return boardrow.ContainsOnly(Column(b, column), s)
	// end::solution[]
	// iftask: return false
}

// DiagRightContainsOnly prüft, ob die Diagonale von links
// oben nach rechts unten nur aus dem übergebenen String besteht.
func DiagRightContainsOnly(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion DiagRight.
	   Verwenden Sie die Funktion boardrow.ContainsOnly.
	*/
	// tag::solution[]
	return boardrow.ContainsOnly(DiagRight(b), s)
	// end::solution[]
	// iftask: return false
}

// DiagLeftContainsOnly prüft, ob die Diagonale von rechts
// oben nach links unten nur aus dem übergebenen String besteht.
func DiagLeftContainsOnly(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion DiagLeft.
	   Verwenden Sie die Funktion boardrow.ContainsOnly.
	*/
	// tag::solution[]
	return boardrow.ContainsOnly(DiagLeft(b), s)
	// end::solution[]
	// iftask: return false
}

// ContainsOnly prüft, ob das Spielbrett nur aus dem übergebenen String besteht.
func ContainsOnly(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie die Funktion boardrow.ContainsOnly.
	*/
	// tag::solution[]
	for row := range b {
		if !RowContainsOnly(b, row, s) {
			return false
		}
	}
	// end::solution[]
	return true
}

// ContainsNone prüft, ob das Spielbrett den übergebenen String nicht enthält.
// D.h. es darf kein Feld auf dem Spielbrett den übergebenen String enthalten.
func ContainsNone(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie eine innere for-Schleife, um die Spalten zu durchlaufen.
	*/
	// tag::solution[]
	for _, row := range b {
		for _, field := range row {
			if field == s {
				return false
			}
		}
	}
	// end::solution[]
	return true
}
