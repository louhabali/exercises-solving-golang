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
	var p []string
	prx := "x = "
	pry := ", y = "
	p = append(p, prx, pry)
	for i, char := range p {
		for _, v := range char {
			z01.PrintRune(v)
		}
		if i == 0 {
			PrintNbr(points.x)
		} else {
			PrintNbr(points.y)
		}
	}
	z01.PrintRune('\n')
}
