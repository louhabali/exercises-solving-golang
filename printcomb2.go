package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	end := false
	for i := '0'; i <= '8'; i++ {
		for j := i + 1; j <= '9'; j++ {

			if end {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			end = true
			z01.PrintRune(i)
			z01.PrintRune(j)
		}
	}
	z01.PrintRune('\n')
}
