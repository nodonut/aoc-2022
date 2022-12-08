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
	var pairs int

	for scnr.Scan() {
		line := scnr.Text()

		firstSplit := strings.Split(line, ",")

		rangeColOne := strings.Split(firstSplit[0], "-")
		rangeColTwo := strings.Split(firstSplit[1], "-")

		min2, err := strconv.Atoi(rangeColTwo[0])
		check(err)
		max2, err := strconv.Atoi(rangeColTwo[1])
		check(err)
		min1, err := strconv.Atoi(rangeColOne[0])
		check(err)
		max1, err := strconv.Atoi(rangeColOne[1])
		check(err)

		fmt.Println(line)
		// if (min1 >= min2 && max1 <= max2) || (min2 >= min1 && max2 <= max1) {
		// 	pairs++
		// }
		if min1 <= max2 && max1 >= min2 {
			pairs++
		}

	}

	fmt.Println(pairs)

}
