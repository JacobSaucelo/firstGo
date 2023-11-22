package main

import "fmt"

func main() {

	hwrro := [3]string{"shero", "lucky", "georgie"}
	fmt.Println(hwrro)
	hewwo := []byte{4, 2, 0}
	fmt.Println(hewwo)
	hewwo[2] = 68
	fmt.Println(hewwo[2])

	hewwo = append(hewwo, 20)
	fmt.Println(hewwo[:3])
	fmt.Println(hewwo)

}

// err aksually main have been redeclared blah blah blah 🤓
