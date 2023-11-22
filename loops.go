package main

import "fmt"

func main() {
	digitsBruh := [3]byte{0, 34, 244}

	// x := 1
	// for x <= 100 {
	// 	fmt.Printf("loading: %v \n", x)
	// 	x++
	// }

	// for i := 0; i < len(digitsBruh); i++ {
	// 	fmt.Println("sheesh: ", digitsBruh[i])
	// }

	for index, value := range digitsBruh {
		fmt.Printf("index: %v value: %v \n", index, value)
	}
	for _, value := range digitsBruh {
		fmt.Printf("value: %v \n", value)
	}

	if true && true {
		fmt.Printf("yo we major")
	} else {
		fmt.Printf("sadge")

	}

}
