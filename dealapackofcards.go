package main

import "fmt"

func DealAPackOfCards (deck []int) {
	deckdividedby4 := len(deck)/4

	for i:= 0; i < 4; i++{
		fmt.Printf("player %d: ", i+1)
		for j:=0; j < deckdividedby4;j++ {
			fmt.Printf("%d", deck[i*deckdividedby4 + j])
			if j < deckdividedby4-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}
func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
// $ go run . | cat -e
// Player 1: 1, 2, 3$
// Player 2: 4, 5, 6$
// Player 3: 7, 8, 9$
// Player 4: 10, 11, 12$
// $