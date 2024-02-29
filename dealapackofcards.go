package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	deckdividedby4 := 3

	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < deckdividedby4; j++ {
			fmt.Printf("%d", deck[i*deckdividedby4+j])
			if j < deckdividedby4-1 {
				fmt.Printf(", ")
			}
		}
		z01.PrintRune('\n')
	}
}
