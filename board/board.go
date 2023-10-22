package board

import "github.com/rhcourses-go/tictactoe/boardrow"

// Board ist ein Spielbrett.
// Wir verwenden eine Slice von BoardRows, um das Spielbrett zu repräsentieren.
type Board []boardrow.BoardRow

// New erzeugt ein Spielbrett.
// Die Funktion erwartet die Breite und Höhe des Spielbretts
// und den String, mit dem das Spielbrett gefüllt werden soll.
func New(width, height int, fill string) Board {
	board := make(Board, height)
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice mit den gewünschten Zeilen zu füllen.
	   Erzeugen Sie für jede Zeile eine neue Slice mit der Funktion boardrow.New.
	*/
	// tag::solution[]
	for i := range board {
		board[i] = boardrow.New(width, fill)
	}
	// end::solution[]
	return board
}

// NewEmpty erzeugt ein leeres Spielbrett.
// Die Funktion erwartet die Breite und Höhe des Spielbretts
// und füllt das Spielbrett mit Leerzeichen.
func NewEmpty(width, height int) Board {
	/* Hinweis:
	   Verwenden Sie die Funktion New.
	*/
	// tag::solution[]
	return New(width, height, " ")
	// end::solution[]
	// iftask: return Board{}
}

// NewFromStrings erzeugt ein Spielbrett aus einer Liste von Strings.
// Die Funktion erwartet eine Liste von Strings, die jeweils eine Zeile des
// Spielbretts repräsentieren.
func NewFromStrings(rows ...string) Board {
	/* Hinweis:
	   Verwenden Sie die Funktion boardrow.NewFromString, um die Zeilen zu erzeugen.
	   Verwenden Sie die Funktion append, um die Zeilen dem Spielbrett hinzuzufügen.
	*/
	// tag::solution[]
	board := Board{}
	for _, row := range rows {
		board = append(board, boardrow.NewFromString(row))
	}
	// end::solution[]
	return board
}

// Row gibt die Zeile an der übergebenen Position zurück.
func Row(b Board, y int) boardrow.BoardRow {
	/* Hinweis:
	   Verwenden Sie die Index-Notation, um die Zeile zu erhalten.
	*/
	// tag::solution[]
	return b[y]
	// end::solution[]
	// iftask: return b[0]
}

// Column gibt die Spalte an der übergebenen Position zurück.
func Column(b Board, x int) boardrow.BoardRow {
	column := make(boardrow.BoardRow, len(b))
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie die Index-Notation, um jeweils das x-te Element der Zeile zu erhalten.
	*/
	// tag::solution[]
	for i, row := range b {
		column[i] = row[x]
	}
	// end::solution[]
	return column
}

// DiagRight gibt die Diagonale von links oben nach rechts unten zurück.
func DiagRight(b Board) boardrow.BoardRow {
	diag := make(boardrow.BoardRow, len(b))
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Da Zeile 0 oben und Spalte 0 links ist, brauchen Sie die Elemente (i,i),
	   also aus jeder Zeile das i-te Element.
	*/
	// tag::solution[]
	for i, row := range b {
		diag[i] = row[i]
	}
	// end::solution[]
	return diag
}

// DiagLeft gibt die Diagonale von rechts oben nach links unten zurück.
func DiagLeft(b Board) boardrow.BoardRow {
	diag := make(boardrow.BoardRow, len(b))
	/* Hinweis:
	   Gehen Sie analog zu GetDiagRight vor.
	   Überlegen Sie sich, welches Element Sie dieses Mal aus jeder Zeile benötigen.
	*/
	// tag::solution[]
	for i, row := range b {
		diag[i] = row[len(b)-1-i]
	}
	// end::solution[]
	return diag
}
