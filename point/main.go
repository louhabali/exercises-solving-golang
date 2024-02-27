package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

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

func main() {
	points := &point{}

	setPoint(points)
	prx := "x= "
	pry := ", y= "
	for _, char := range prx {
		z01.PrintRune(char)
	}
	PrintNbr(points.x)
	for _, char := range pry {
		z01.PrintRune(char)
	}
	PrintNbr(points.y)
}
