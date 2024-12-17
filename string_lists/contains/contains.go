package contains

// Erwartet einen String und eine Liste von Strings.
// Gibt true zurück, wenn der String in der Liste enthalten ist, ansonsten false.
func Contains(strings []string, search string) bool {

	if len(strings) == 0 { //Wenn die Länge des Strings == 0, sit das element auf keinen Fall enthalten
		return false
	}

	if strings[len(strings)-1] == search { // Wenn das letzte Element das richtige ist soll true returnt werden
		return true
	}

	return Contains(strings[:len(strings)-1], search) //nochmal Wiederholen nur ohne das bereits gecheckte Element von strings[]
}
