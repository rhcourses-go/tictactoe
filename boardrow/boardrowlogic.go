package boardrow

// ContainsOnly prüft, ob die Zeile nur aus dem übergebenen String besteht.
func ContainsOnly(row BoardRow, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice zu durchlaufen.
	   Brechen Sie die Schleife ab, sobald ein Element ungleich dem übergebenen String ist.
	*/
	// tag::solution[]
	for _, v := range row {
		if v != s {
			return false
		}
	}
	return true
	// end::solution[]
	// iftask: return false
}

// IsEmpty prüft, ob die Zeile leer ist.
func IsEmpty(row BoardRow) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion ContainsOnly.
	*/
	// tag::solution[]
	return ContainsOnly(row, " ")
	// end::solution[]
	// iftask: return false
}
