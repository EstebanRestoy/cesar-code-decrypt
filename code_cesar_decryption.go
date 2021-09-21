package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("Pls give the text in arg to the run command")
		os.Exit(3)

	}
	text := string(argsWithoutProg[0])

	for i := 0; i < 26; i++ {
		textDecrypted := ""
		for j := range text {
			textDecrypted += string(nextChar(i, text[j]))
		}
		fmt.Print("Shift: " + strconv.Itoa(i) + " Text : ")
		fmt.Println(textDecrypted)
	}
}

func nextChar(rotate int, ch byte) byte {

	for i := 0; i < rotate; i++ {
		if ch += 1; ch > 'z' {
			ch = 'a'
		}
	}

	return ch
}
