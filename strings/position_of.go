package strings

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {

	//Suchen nach c in s
	for i, el := range s {

		//Fuer alle Elemente Vergleichen ob c das Element ist
		if rune(c) == el {

			//Wenn c das erste mal == el ist, soll die Funktion die Stelle herausgeben
			return i
		}
	}

	//Wenn nicht vorkommt, return len(s)
	return len(s)
}
