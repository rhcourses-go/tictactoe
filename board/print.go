package board

import (
	"fmt"

	"github.com/rhcourses-go/tictactoe/boardrow"
)

// Divider gibt den Trenn-String für die Zeilen des Spielbretts zurück.
// Die Funktion erwartet die Länge der Zeile.
// Der Trenn-String ist von der Form "+---+---+---+".
func RowDivider(length int) string {
	result := "+"
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um den Trenn-String zu erzeugen.
	   Hängen Sie in jedem Schleifendurchlauf den String "---+" an.
	*/
	// tag::solution[]
	for i := 0; i < length; i++ {
		result += "---+"
	}
	// end::solution[]
	return result
}

// PrintBoard gibt das Spielbrett auf die Konsole aus.
// Die Funktion erwartet das Spielbrett.
func PrintBoard(b Board) {
	divider := RowDivider(len(b[0]))
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen des Spielbretts zu durchlaufen.
	   Verwenden Sie die Funktion boardrow.Print, um die Zeile auszugeben.
	   Geben Sie nach jeder Zeile den Trenn-String aus.
	*/
	// tag::solution[]
	fmt.Println(divider)
	for _, row := range b {
		boardrow.Print(row)
		fmt.Println(divider)
	}
	// end::solution[]
}
