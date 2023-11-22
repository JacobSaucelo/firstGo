package main

import (
	"fmt"
	"strconv"
)

var fard []int = []int{122, 545, 76, 123, 65}

func main() {
	// greetUser("tae")
	// goodBye("asdfBobo")
	shitter(fard, goodBye)
	fmt.Println(num2string(2000), "dolllers")
}

func greetUser(name string) {
	fmt.Printf("henlo!! ^_^ %v \n", name)
}

func goodBye(name string) {
	fmt.Printf("byee!! UwU :3 lmao *starts twerking* %v \n", name)
}

func shitter(container []int, bye func(string)) {
	for _, value := range container {
		bye("bobo!")
		fmt.Printf("you owe me $%v \n", value)
	}
}

func num2string(num int) string {
	return strconv.Itoa(num)
}
