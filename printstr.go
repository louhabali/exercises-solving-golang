package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for x := 0; x < len(s); x++ {
		z01.PrintRune(s[x])
	}
}

