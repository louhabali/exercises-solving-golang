package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for f := '0'; f <= '9'; f++ {
					
					if (i <= '9' && j <= '8' && k <= '9' && f <= '9' ) {
                    z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(f)
					z01.PrintRune(',')
					z01.PrintRune(' ')
					}
				

				}
			}
		}
	}
	z01.PrintRune('\n')

}