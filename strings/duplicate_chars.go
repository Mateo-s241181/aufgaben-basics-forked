package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {

	//Definieren von einem neuen String (zunaechst leer)
	duplicatedString := ""

	//Auslesen der Einzelnen Symbole aus s

	for _, el := range s {

		//Speichern des Symbols zweimal in duplicatedString
		duplicatedString = duplicatedString + string(el) + string(el)
	}

	return duplicatedString
}
