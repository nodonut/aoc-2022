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

	myScanner := bufio.NewScanner(file)
	var totalMap map[string]int
	var total int

	var count int = 0
	firstThreeLines := ""

	for myScanner.Scan() {
		line := myScanner.Text()

		// firstHalf, secondHalf := partition(line)

		firstThreeLines += line + "+"
		count++

		if count == 3 {
			splitStr := strings.Split(firstThreeLines, "+")
			_, same := compare(splitStr[0], splitStr[1])
			totalMap, _ = compare(splitStr[2], same)
			for _, v := range totalMap {
				total += v
			}
			count = 0
			firstThreeLines = ""
		}

	}

	fmt.Println(total)
}

func partition(s string) (string, string) {
	pivot := len(s) / 2

	return s[:pivot], s[pivot:]
}

func compare(s string, sb string) (map[string]int, string) {
	totalMap := make(map[string]int)
	same := ""

	for _, char := range sb {
		if strings.Contains(s, string(char)) {
			same += string(char)
			totalMap[string(char)] = int(getCharScore(char))
		}
	}

	return totalMap, same
}

func getCharScore(c rune) int32 {
	var lt rune

	if c >= 'a' && c <= 'z' {
		lt = c - 'a' + 1

	} else {
		lt = c - 'A' + 26 + 1
	}

	return lt
}
