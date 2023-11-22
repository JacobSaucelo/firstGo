package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	message := "haii!! ^_^ hii! :3"
	// message2 := "Yo sup G!"
	profanity := "bobo whahahaha"

	fmt.Println(len(message))
	fmt.Println(strings.Contains(message, "haiii"))
	fmt.Println(strings.Index(message, "hii"))
	fmt.Println(strings.ReplaceAll(profanity, "bobo", "****"))
	fmt.Println(strings.Split(message, " "))

	scores := []int{45, 12, 65, 89}
	sort.Ints(scores)
	fmt.Println(scores)
}
