package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./prod.txt")
	check(err)
	defer file.Close()

	scnr := bufio.NewScanner(file)
	characters := ""

	for scnr.Scan() {
		line := scnr.Text()

		for i, l := range line {
			if strings.Index(characters, string(l)) >= 0 && len(characters) < 14 {
				characters = characters[strings.Index(characters, string(l))+1:]
				characters += string(l)
			} else if len(characters) == 14 {
				fmt.Println("Start of marker: ", i)
				characters = ""
				break
			} else {
				characters += string(l)
			}
		}
	}
}
