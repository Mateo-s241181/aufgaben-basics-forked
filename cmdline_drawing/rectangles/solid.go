package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	for i := 0; i < height; i++ {
		Drawline(width, "#")
	}
}

func Drawline(width int, symbol string) {
	for i := 0; i < width; i++ {
		fmt.Print(symbol)
	}
	fmt.Println("")
}
