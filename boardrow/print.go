package boardrow

import "fmt"

// Print gibt die Zeile des Spielbretts auf die Konsole aus.
// Die Funktion erwartet eine Zeile des Spielbretts.
func Print(row BoardRow) {
	output := "| "
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice zu durchlaufen.
	   Hängen Sie an die Variable output jeweils das Element der Slice
	   und den String "|" an.
	*/
	// tag::solution[]
	for _, v := range row {
		output += v + " | "
	}
	output = output[:len(output)-1]
	// end::solution[]
	fmt.Println(output)
}
