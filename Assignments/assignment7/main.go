package main

import (
	"fmt"
	"math/rand/v2"
)

// Create a program that rolls two dice (1 to 6) fifty times.
// Display the number rolls and the outcomes in sequential order.Resulting rolls are to be processed in the following manner:
// [Random Numbers][Switches]
// 7 and 11 are to be called NATURAL
// 2 is called SNAKE-EYES-CRAPS
// 3 and 12 is called LOSS-CRAPS
// Any other combination is called NEUTRAL.
func main() {
	for i := 0; i < 30; i++ {
		first, second := roll2Dice()
		fmt.Printf("got %d and %d\n", first, second)
		fmt.Printf("The result is %s\n", switchCase(first, second))
	}
}

func roll2Dice() (int, int) {
	return randRange(1, 7), randRange(1, 6)
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func switchCase(first, second int) string {
	sum := first + second
	switch sum {
	case 11, 7:
		return "NATURAL"
	case 2:
		return "SNAKE-EYES-CRAPS"
	case 12, 3:
		return "LOSS-CRAPS"
	default:
		return "NEUTRAL"
	}
}
