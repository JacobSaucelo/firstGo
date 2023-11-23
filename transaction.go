package main

import "fmt"

func newTransaction(buyer string) transactionType {
	return transactionType{
		buyerName: buyer,
		item:      map[string]shopItems{},
		total:     0,
	}
}

func (t *transactionType) buyItem(item shopItems) {
	t.item[item.name] = item
}

func (t *transactionType) showTransaction() {

	for key, values := range t.item {
		fmt.Printf("%20v: %v gold | rarity %v\n", key, values.price, values.rarity)
		t.total += values.price
	}
	fmt.Println("_____________________________________________")
	fmt.Printf("total: %20v gold\n", t.total)
}
