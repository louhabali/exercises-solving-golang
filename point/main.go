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

func main() {
	points := &point{}

	setPoint(points)

	str := "x = " + string('0'+(points.x/10)) + string('0'+(points.x%10)) + ", y = " + string('0'+(points.y/10)) + string('0'+(points.y%10)) + "\n"
	for _, character := range str {
		z01.PrintRune(character)
	}
}
