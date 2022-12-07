package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	reversed := ""
	moves := ""

	stack := [][]string{}

	for scnr.Split(bufio.ScanLines); scnr.Scan(); {
		line := scnr.Text()

		if !strings.Contains(line, "move") && !strings.Contains(line, "1") {
			reversed = line + "\n" + reversed
		}

		if strings.Contains(line, "move") {
			moves = moves + "\n" + line
		}
	}

	revScnr := bufio.NewScanner(strings.NewReader(reversed))

	for revScnr.Scan() {
		line := revScnr.Text()
		crate := walkTheLine(line)
		if len(crate) > 0 {
			stack = append(stack, crate)
		}
	}

	matrix := convertToTwoDArray(stack)

	moveScnr := bufio.NewScanner(strings.NewReader(moves))

	for moveScnr.Scan() {
		line := moveScnr.Text()

		if len(line) > 0 {
			length, err := strconv.Atoi(strings.Split(line, " ")[1])
			check(err)
			from, err := strconv.Atoi(strings.Split(line, " ")[3])
			check(err)
			to, err := strconv.Atoi(strings.Split(line, " ")[5])
			check(err)

			move(matrix, length, from, to)
		}
	}
	printResult(matrix)
}

func walkTheLine(l string) []string {
	crates := []string{}

	for i := 0; i < len(l); i++ {
		crate := fmt.Sprintf("%c%c%c", rune(l[i]), rune(l[i+1]), rune(l[i+2]))
		crates = append(crates, crate)

		i += 3
	}

	return crates
}

func convertToTwoDArray(s [][]string) [][]string {
	stack := [][]string{}

	for i := 0; i <= len(s); i++ {
		level := []string{}
		for j := 0; j < len(s); j++ {
			if s[j][i] != "   " {
				level = append(level, s[j][i])
			}
		}
		stack = append(stack, level)
	}

	return stack
}

func move(m [][]string, l int, f int, t int) {
	from := f - 1
	to := t - 1
	fromLen := len(m[from])
	toItems := m[to][:]
	var fromItems []string

	for i := l; i > 0; i-- {
		if fromLen-i == 0 {
			fromItems = []string{}
		} else {
			fromItems = m[from][:fromLen-l]

		}
		fromItem := m[from][fromLen-i]
		toItems = append(toItems, fromItem)
	}

	m[from] = fromItems
	m[to] = toItems
}

func printResult(m [][]string) {
	firstLength := len(m)

	result := ""

	for i := 0; i < firstLength; i++ {
		fmt.Println(len(m[i]))
		result = result + string(m[i][len(m[i])-1][1])
	}

	fmt.Println(result)
}
