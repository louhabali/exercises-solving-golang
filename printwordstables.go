package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, ch := range a {
		s := ch
		for _, f := range s {
			z01.PrintRune(f)
		}
		z01.PrintRune('\n')
	}
}
