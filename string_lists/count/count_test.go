package count

import "fmt"

func ExampleCount() {
	s1 := []string{"a", "b", "a", "c", "b", "a"}

	counter := 0

	fmt.Println(Count(s1, "a", counter))
	fmt.Println(Count(s1, "b", counter))
	fmt.Println(Count(s1, "c", counter))

	// Output:
	// 3
	// 2
	// 1
}
