package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	name, _ := greetName(reader)

	fmt.Println("hello", name)

}

func greetName(reader *bufio.Reader) (string, error) {

	fmt.Print("what is your name: ")

	name, err := reader.ReadString('\n')

	return name, err

}
