package main

import "fmt"

var commonItems2 = map[byte]shopItems{
	1: {
		id:     1,
		name:   "Garlic",
		rarity: 1,
		price:  2,
	},
	2: {
		id:     2,
		name:   "Leather Bag",
		rarity: 1,
		price:  2,
	},
	3: {
		id:     3,
		name:   "Banana",
		rarity: 1,
		price:  3,
	},
	4: {
		id:     4,
		name:   "Wooden Buckler",
		rarity: 1,
		price:  4,
	},
}

func main() {
	fmt.Println("added ", createItem("Broom", 1, 4))
	fmt.Println("___________________________")

	for _, value := range commonItems2 {
		fmt.Printf("%v: %v gold | rarity %v\n", value.name, value.price, value.rarity)
	}

}

func createItem(n string, r byte, p byte) shopItems {
	item := shopItems{
		id:     byte(len(commonItems2) + 1),
		name:   n,
		rarity: r,
		price:  p,
	}

	commonItems2[byte(len(commonItems2)+1)] = item

	return item
}

//? shop.go - types.go
