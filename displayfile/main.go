package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	contenu, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("the error you made is : ", err)
		return
	}
	fmt.Print(string(contenu))
}