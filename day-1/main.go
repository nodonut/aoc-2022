package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

	var count int64

	total := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if num, err := strconv.ParseInt(line, 10, 64); err == nil {
			count += num
		} else {
			total = append(total, int(count))
			count = 0
		}
	}

	sort.Ints(total)
	sort.Sort(sort.Reverse(sort.IntSlice(total)))

	topThree := 0

	for _, v := range total[0:3] {
		topThree += v
	}

	fmt.Println(topThree)

	check(scanner.Err())
}
