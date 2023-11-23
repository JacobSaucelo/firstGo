package main

import "fmt"

func main() {
	commonItems := map[string]byte{
		"Broom":          4,
		"Garlic":         2,
		"Leather Bag":    2,
		"Banana":         3,
		"Wooden Buckler": 4,
	}

	fmt.Println(commonItems)

	for index, value := range commonItems {
		fmt.Printf("%v: %v gold \n", index, value)
	}
}
