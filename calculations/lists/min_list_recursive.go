package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.

func MinListRecursive(nums []int) int {
	if len(nums) == 0 { //Länge 0 -> 0
		return 0
	}

	minNum := nums[len(nums)-1] //letztes element wird als kleinstes gesetzt

	if len(nums) == 1 { //Länge 1 -> Element im slice ist das kleinste
		return minNum
	}

	if nums[len(nums)-2] < minNum { //Wenn das Vorletzte Element kleiner als das letzte ist, ist es das neue minNum und lösch das letzte element aus dem slice
		minNum = nums[len(nums)-2]
		nums = nums[:len(nums)-1]
	}

	if len(nums) == 1 { //Muss ein zweites mal Testen ob die Länge des Slices = 1 ist, damit in der nächten Funktion nicht der Index [-1] verwendet wird
		return minNum
	}

	if nums[len(nums)-2] >= minNum { //Wenn das vorletzte Element größer als das letzte ist, lösch das vorletzte element
		nums = nums[:len(nums)-2]
		nums = append(nums, minNum)
	}

	return MinListRecursive(nums)
}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
