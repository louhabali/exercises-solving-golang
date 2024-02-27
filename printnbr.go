package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(nb int) {
	n := '0'
	subn := '0'
	for i := 0; i < nb/10; i++ {
		n++
	}
	z01.PrintRune(n)
	for i := 0; i < nb%10; i++ {
		subn++
	}
	z01.PrintRune(subn)
}
