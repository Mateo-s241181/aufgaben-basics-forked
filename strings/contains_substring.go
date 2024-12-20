package strings

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(str, sub string) bool {

	//Laenge von Unveraendertem Substring in Variable speichern
	length := len(sub)
	counter := 0

	if len(sub) == 0 {
		return true
	}

	return InnerContainsSubstring(str, sub, length, counter)
}

func InnerContainsSubstring(s, t string, l, count int) bool {

	//Wenn der counter nach einigen Iterationen die laenge 0 erreicht hat, ist t nicht in s
	if len(t) == 0 {
		return false
	}

	//letztes element des strings mit dem letzten des substrings verleichen
	if s[len(s)-1] == t[len(t)-1] {
		count++

		// Wenn der counter die selbe laenge wie t erreicht hat, ist t vollstaendig in s enthalten
		if count == l {
			return true
		}

		//Die jeweils letzten elemente rausloeschen
		return InnerContainsSubstring(s[:len(s)-1], t[:len(t)-1], l, count)

	} else {
		count = 0

		//Substring ungekuerzt lassen
		return InnerContainsSubstring(s[:len(s)-1], t, l, count)
	}

}
