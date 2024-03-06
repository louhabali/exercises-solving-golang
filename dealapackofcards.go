package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	d := 3
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < d; j++ {
			fmt.Printf("%d", deck[i*d+j])
			if j < d-1 {
				fmt.Printf(", ")
			}
		}
		z01.PrintRune('\n')
	}
}
