package strings

import (
	"slices"
	"strings"
)

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	//Beide strings werden in Slices umgewandelt

	s1_slice := []byte(s1)
	s2_slice := []byte(s2)

	//Wenn die Länge der Listen ungleich sind
	if len(s1_slice) != len(s2_slice) {
		return false
	}

	//Slices werden geordnet
	slices.Sort(s1_slice)
	slices.Sort(s2_slice)

	//Falls Slices gleich sind sind die Strings Anagramme
	return string(s1_slice) == string(s2_slice)
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	//Beide strings werden in Slices umgewandelt

	s1_slice := []byte(strings.ToLower(s1))
	s2_slice := []byte(strings.ToLower(s2))

	//Wenn die Länge der Listen ungleich sind
	if len(s1_slice) != len(s2_slice) {
		return false
	}

	//Slices werden geordnet
	slices.Sort(s1_slice)
	slices.Sort(s2_slice)

	//Falls Slices gleich sind sind die Strings Anagramme
	return string(s1_slice) == string(s2_slice)
}
