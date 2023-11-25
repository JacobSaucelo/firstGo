package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// transaction1 := newTransaction("jacob")

	// transaction1.buyItem(shopItems{
	// 	id:     10,
	// 	rarity: 1,
	// 	name:   "stone",
	// 	price:  1,
	// })
	// transaction1.buyItem(shopItems{
	// 	id:     12,
	// 	rarity: 3,
	// 	name:   "dagger",
	// 	price:  3,
	// })
	// transaction1.buyItem(shopItems{
	// 	id:     1,
	// 	rarity: 2,
	// 	name:   "broom",
	// 	price:  4,
	// })

	// transaction1.showTransaction()
}

func showShop(buyer string, reader *bufio.Reader) (string, error) {
	options, err := reader.ReadString('\n')

	return strings.TrimSpace(options), err
}

// go run market.go types.go transaction.go
