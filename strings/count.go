package strings

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {
	//counter ist am anfang 0
	counter := 0

	//for loop geht durch den string und schaut ob c vorkommt
	for _, char := range s {
		if char == c {
			//Wenn c vorkommt geht der counter eins nach oben
			counter++
		}
	}

	return counter
}

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {

	result := 0
	for _, char := range s {
		if char == 'A' {
			result++
		}
	}
	return result

}
