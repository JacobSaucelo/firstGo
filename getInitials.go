package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getInitials("hewwo jacob"))
	fmt.Println(getInitials("jacob saucelo"))
	fmt.Println(getInitials("jacob"))
}

func getInitials(name string) (string, string) {
	initials := strings.Split(strings.ToUpper(name), " ")

	for index, value := range initials {
		initials[index] = value[:1]
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""
}
