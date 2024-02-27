package main

import (
	"piscine"

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

func main() {
	points := &point{}

	setPoint(points)
	prx := "x = "
	pry := ", y = "
	for _, char := range prx {
		z01.PrintRune(char)
	}
	piscine.PrintNbr(points.x)
	for _, char := range pry {
		z01.PrintRune(char)
	}
	piscine.PrintNbr(points.y)
	z01.PrintRune('\n')
}
