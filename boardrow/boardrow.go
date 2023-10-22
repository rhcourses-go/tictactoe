package boardrow

// BoardRow ist eine Zeile des Spielbretts.
// Wir verwenden Strings, um die Spielfiguren zu repräsentieren.
// Daher ist BoardRow ein Slice von Strings.
type BoardRow []string

// New erzeugt eine Zeile des Spielbretts.
// Die Funktion erwartet die Länge der Zeile und den String,
// mit dem die Zeile gefüllt werden soll.
func New(length int, fill string) BoardRow {
	row := make(BoardRow, length)
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice mit dem gewünschten String zu füllen.
	*/
	// tag::solution[]
	for i := range row {
		row[i] = fill
	}
	// end::solution[]
	return row
}

// NewEmpty erzeugt eine leere Zeile des Spielbretts.
// Die Funktion erwartet die Länge der Zeile
// und füllt die Zeile mit Leerzeichen.
func NewEmpty(length int) BoardRow {
	/* Hinweis:
	   Verwenden Sie die Funktion MakeBoardRow.
	*/
	// tag::solution[]
	return New(length, " ")
	// end::solution[]
	// iftask: return BoardRow{}
}

// NewFromString erzeugt eine Zeile des Spielbretts aus einem String.
// Jedes Zeichen im String repräsentiert ein Feld der Zeile.
func NewFromString(row string) BoardRow {
	/* Hinweis:
	   Verwenden Sie die Funktion NewEmpty.
	   Iterieren Sie anschließend über den String und setzen Sie die einzelnen Zeichen.
	*/
	// tag::solution[]
	boardRow := NewEmpty(len(row))
	for i, c := range row {
		boardRow[i] = string(c)
	}
	return boardRow
	// end::solution[]
	// iftask: return BoardRow{}
}
