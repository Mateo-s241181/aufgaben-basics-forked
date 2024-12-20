package strings

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {

	if len(s) == 0 {
		return false
	}

	if s[len(s)-1] == c {
		return true
	}

	return Contains(s[:len(s)-1], c)
}
