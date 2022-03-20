package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you can open")
	case 3:
		fmt.Println("Dice value is 3 and you can open")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 and you can open")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 and you can open")
	case 6:
		fmt.Println("Dice value is 6 and you can open")
	default:
		fmt.Println("what was that")
	}
}
