package main

import "fmt"

func main() {
	var gold byte = 1

	commonItems := map[string]byte{
		"Broom":          4,
		"Garlic":         2,
		"Leather Bag":    2,
		"Banana":         3,
		"Wooden Buckler": 4,
	}

	buyItem(commonItems["Broom"], &gold)

	// fmt.Println(commonItems)

	// for index, value := range commonItems {
	// 	fmt.Printf("%v: %v gold \n", index, value)
	// }
}

func buyItem(item byte, currentGold *byte) byte {
	if *currentGold < item {
		fmt.Println("Insufficient funds.")
		return 0
	}
	fmt.Println("transaction successful.")
	return *currentGold - item
}

// * - ptr value
// & - memory address value
