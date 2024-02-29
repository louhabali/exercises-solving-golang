package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("player %d: ", i+1)
		for j := 0; j < 3; j++ {
			if j < 2 {
				fmt.Printf("%d, ", deck[i*3+j])
			} else {
				fmt.Printf("%d", deck[i*3+j])
			}
		}
		z01.PrintRune('\n')
	}
}
