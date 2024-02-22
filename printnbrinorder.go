package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var array []int

	if n == 0 {
		z01.PrintRune(48)
	}

	for n > 0 {
		array = append(array, n%10)
		n /= 10
	}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				c := array[i]
				array[i] = array[j]
				array[j] = c
			}
		}
	}
	for k := 0; k < len(array); k++ {
		z01.PrintRune(rune(array[k] + 48))
	}
}
