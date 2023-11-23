package main

type shopItems struct {
	id     byte
	rarity byte
	name   string
	price  byte
}

type transactionType struct {
	buyerName string
	item      map[string]shopItems
	total     byte
}
