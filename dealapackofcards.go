package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	deckdividedby4 := 3

	for i := 0; i < 4; i++ {
		fmt.Printf("player %d: ", i+1)
		for j := 0; j < deckdividedby4; j++ {
			fmt.Printf("%d", deck[i*deckdividedby4+j])
			if j < deckdividedby4-1 {
				fmt.Print(", ")
			}
		}
		z01.PrintRune('\n')
	}
}
