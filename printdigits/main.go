package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "0123456789"
	for _, ch := range alphabet {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
