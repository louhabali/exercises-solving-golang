package main

import (
	"fmt"
	"os"
)

func main() {
	osarr := os.Args[1:]
	for _, ch := range osarr {
		if ch == "01" || ch == "galaxy" || ch == "galaxy 01" {
			fmt.Println("Alert!!!")
		} else {
			fmt.Print("")
		}
	}
}
