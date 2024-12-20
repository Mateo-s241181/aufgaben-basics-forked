package make_unique

import "strconv"

// Erwartet eine Liste von Strings.
// Hängt Zahlen an alle mehrfach vorkommenden Strings an, um sie eindeutig zu machen.

func MakeUnique(strings []string) []string {

	for i, el1 := range strings {

		//Der Counter soll zunächst 1 sein
		counter := 1

		for j, el2 := range strings {

			if strings[i] == el2 && i != j {
				//Counter erhöht sich
				counter++
				//Das Element an der Stelle j wird das Suffix _counter angehängt
				strings[j] = el1 + "_" + strconv.Itoa(counter)
			}
		}
	}

	//Strings wird returned, nachdem die Suffixe angehänt wurden
	return strings
}
