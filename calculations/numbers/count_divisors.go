package numbers

import "math"

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {

	numberOfDivisors := 0

	//Wenn n die Zahl eis ist, soll der Code 1 als Anzahl zurückgeben
	if n == 1 {
		return 1
	}

	//Für die restlichen Möglichkeiten wird gecheckt ob die division durch i eine ganze zahl ergibt
	//Es wird nur bis zur Wurzel gecheckt, da in [1, sqrt(n)) genauso viele Teiler sind wie in (sqrt(n), n]
	for i := 1.0; i < math.Sqrt(float64(n)); i++ {

		if float64(n)/i == math.Trunc(float64(n)/i) {
			numberOfDivisors++
		}
	}

	//Sonderfall: Wenn die Sqrt(n) eine ganze Zahl ist, dann ist es auch ein Teiler von n und muss also betrachtet werden ( +1 )
	if math.Sqrt(float64(n)) == math.Trunc(math.Sqrt(float64(n))) {
		return (numberOfDivisors * 2) + 1
	}

	return numberOfDivisors * 2
}
