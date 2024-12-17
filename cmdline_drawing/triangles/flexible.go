package triangles

import (
	"fmt"
)

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {
	for i := 1; i <= length; i++ {

		if i == 1 || i == length {

			for k := 0; k < i; k++ {
				fmt.Print(outer)
			}

		} else {

			for j := 1; j <= i; j++ {

				if j == 1 || j == i {

					fmt.Print(outer)

				} else {

					fmt.Print(inner)
				}
			}
		}

		fmt.Println("")

	}
}
