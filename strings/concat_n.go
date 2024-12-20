package strings

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {

	newString := s

	for i := 0; i < n-1; i++ {
		newString = newString + sep + s
	}

	return newString
}
