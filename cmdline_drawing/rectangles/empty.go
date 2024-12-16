package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	for i := 0; i < height; i++ {
		if i == 0 || i == height-1 {
			for i := 0; i < width; i++ {
				fmt.Printf("#")
			}
		} else {
			for i := 0; i < width; i++ {
				if i == 0 || i == width-1 {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
