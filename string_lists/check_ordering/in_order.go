package check_ordering

// Erwartet eine Liste von Strings und zwei einzelne Strings.
// Überprüft, ob die beiden Strings in der Liste in der gegebenen Reihenfolge vorkommen.
// Gibt `true` zurück, wenn das der Fall ist, ansonsten `false`.
func CheckOrdering(strings []string, first, second string) bool {

	//Wenn der string unter 2 elemente enthaelt und in der letzten iteration noch nicht true retured worden ist, sind die Elemente nicht in der Richtigen Reihenfolge

	if len(strings) == 1 || len(strings) == 0 {
		return false
	}

	//Wenn die beiden Strings direkt hintereinander in der richtigen Reihenfolge vorkommen soll true returned werden

	if strings[len(strings)-2] == first && strings[len(strings)-1] == second {
		return true
	}

	// Wenn das letzte Element von strings[] == second ist, und das vorletzte nicht first ist, soll das vorletzte mit dem letzten ueberschrieben werden

	if strings[len(strings)-1] == second && strings[len(strings)-2] != first {
		strings[len(strings)-2] = strings[len(strings)-1]
	}

	//In der naechsten Iteration soll das letzte Element nicht mehr betrachtet werden. So wird der String verkleinert und die zu betrachtenden Elemente eingeschraenkt

	return CheckOrdering(strings[:len(strings)-1], first, second)
}

// REMARKS
// - Diese Aufgabe ist eine komplexere Variante der Aufgabe "Prüfen, ob ein Element in einer Liste vorkommt".
// - Sie können die Lösung der einfachen Variante als Grundlage verwenden und diese entsprechend erweitern.
