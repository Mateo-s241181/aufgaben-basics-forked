package count

// Erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Liefer die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string, counter int) int {

	if len(strings) == 0 { //Wenn die Länge des Strings == 0, soll der counter returnt werden
		return counter
	}

	if strings[len(strings)-1] == search { // Wenn das letzte Element das richtige ist soll der counter um eins erhoeht weerden
		counter++
	}

	return Count(strings[:len(strings)-1], search, counter) //nochmal Wiederholen nur ohne das bereits gecheckte Element von strings[]
}
