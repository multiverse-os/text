package main

import (
	"fmt"
	"math/rand"
	"time"

	dice "github.com/multiverse-os/symbols/games/dice"
)

func randomNumber(maxNumber int) int {
	randomNumber := rand.Intn(maxNumber) + 1
	return randomNumber
}

func drawDice(number int) string {
	switch number {
	case 1:
		return dice.Symbols["one"]
	case 2:
		return dice.Symbols["two"]
	case 3:
		return dice.Symbols["three"]
	case 4:
		return dice.Symbols["four"]
	case 5:
		return dice.Symbols["five"]
	default: // 6
		return dice.Symbols["six"]
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("dice-cli\n")
	fmt.Printf("========\n")

	diceOne := randomNumber(6)
	diceTwo := randomNumber(6)

	fmt.Printf("=> %s + %s = %v\n", drawDice(diceOne), drawDice(diceTwo), (diceOne + diceTwo))

}
