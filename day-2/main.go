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
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scoreMap := map[string]int{
	// 	"X": 1,
	// 	"Y": 2,
	// 	"Z": 3,
	// }
	opScoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	var total int

	for scanner.Scan() {
		line := scanner.Text()
		var lineScore int = 0

		if strings.Contains(line, "A") {
			if strings.Contains(line, "X") {
				lineScore += opScoreMap["C"]
				lineScore += 0
			}
			if strings.Contains(line, "Y") {
				lineScore += opScoreMap["A"]
				lineScore += 3
			}
			if strings.Contains(line, "Z") {
				lineScore += opScoreMap["B"]
				lineScore += 6
			}
		} else if strings.Contains(line, "B") {
			if strings.Contains(line, "X") {
				lineScore += opScoreMap["A"]
				lineScore += 0
			}
			if strings.Contains(line, "Y") {
				lineScore += opScoreMap["B"]
				lineScore += 3
			}
			if strings.Contains(line, "Z") {
				lineScore += opScoreMap["C"]
				lineScore += 6
			}
		} else if strings.Contains(line, "C") {
			if strings.Contains(line, "X") {
				lineScore += opScoreMap["B"]
				lineScore += 0
			}
			if strings.Contains(line, "Y") {
				lineScore += opScoreMap["C"]
				lineScore += 3
			}
			if strings.Contains(line, "Z") {
				lineScore += opScoreMap["A"]
				lineScore += 6
			}
		}
		total += lineScore
	}

	fmt.Printf("Total: %d\n", total)

	check(scanner.Err())
}
