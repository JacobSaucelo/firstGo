package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var storeItems = []string{
	"Stun grenade",
	"Boombox",
	"TZP-Inhalant",
	"Zap gun",
	"Jetpack",
	"Extension ladder",
	"Radar-booster",
}

func main() {
	mainMenu()
}

func mainMenu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Main Menu: \n > shop\n > brightness\n > exit\n")
	pick, _ := reader.ReadString('\n')

	switch strings.TrimSpace(pick) {
	case "shop":
		showStore()
	case "brightness":
		brightness()
		// fmt.Println("you picked brightness.")
	case "exit":
		fmt.Println("you picked exit.")
	default:
		fmt.Println("terminal: Not a valid input.")
		mainMenu()
	}
}

func showStore() {
	reader := bufio.NewReader(os.Stdin)

	for key, values := range storeItems {
		fmt.Printf("%v // %v\n", key+1, values)
	}

	picked, _ := reader.ReadString('\n')

	switch strings.TrimSpace(picked) {
	case "1":
		fmt.Println("you bought", storeItems[0])
	case "2":
		fmt.Println("you bought", storeItems[1])
	case "3":
		fmt.Println("you bought", storeItems[2])
	case "4":
		fmt.Println("you bought", storeItems[3])
	case "5":
		fmt.Println("you bought", storeItems[4])
	case "6":
		fmt.Println("you bought", storeItems[5])
	case "7":
		fmt.Println("you bought", storeItems[6])
	default:
		fmt.Println("terminal: That is not a valid input...")
		showStore()
	}
}

func brightness() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Set brightness to 1.0-10.0")
	value, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("terminal: Invalid input.")
		brightness()
	}

	bValue := strings.TrimSpace(value)
	parsedValue, err := strconv.ParseFloat(bValue, 64)

	if parsedValue > 10 {
		fmt.Println("terminal: Invalid input. up to 1.0-10.0 values only")
		brightness()
	}

	fmt.Println("Brightness is set to", parsedValue)

}
