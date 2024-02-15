package piscine

import "github.com/01-edu/z01"

func IsNegative(num int) {
	if num >= 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
	z01.PrintRune('\n')
}
