package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	for i := 0; i < height; i++ {
		if i == 0 || i == height-1 {
			Drawline(width, outer)
		} else {
			DrawInnerRows(width, inner, outer)
		}
	}
}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.

func DrawInnerRows(width int, fillingSymbol, borderSymbol string) {

	for i := 0; i < width; i++ {

		if i == 0 || i == width-1 {

			fmt.Print(borderSymbol)

		} else {

			fmt.Print(fillingSymbol)
		}
	}
	fmt.Println("")
}
