package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.ReadFile("prod.txt")
	check(err)

	input := strings.Split(string(file), "\n")
	input = input[:len(input)-1] // remove empty last line

	var matrix [][]string
	var edges int

	for i, v := range input {
		if i == 0 || i == len(input)-1 { // first and last are always visible
			edges += len(input)
		} else {
			edges += 2 // only the first and last element for the inbetween are visible
		}
		matrix = append(matrix, strings.Split(v, ""))
	}

	invertedMatrix := invert(matrix) // for top and bottom calculations
	compare(matrix, &edges, &invertedMatrix)
}

func compare(arr [][]string, edges *int, inv *[][]string) {
	result := []int{} // for part two

	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr[i])-1; j++ {
			currentValue, err := strconv.Atoi(arr[i][j])
			check(err)
			left := arr[i][:j]
			countLeft := compareLeftRight(left, currentValue, j, "left")

			if countLeft == 0 {
				right := arr[i][j+1:]
				countRight := compareLeftRight(right, currentValue, j, "right")
				*edges += countRight

				if countRight == 0 {
					vert := compareTopBottom(*inv, currentValue, i, j)
					*edges += vert
				}
			}
			*edges += countLeft

			ptl := partTwo(left, currentValue, j, "left")          // partTwoLeft value
			ptr := partTwo(arr[i][j+1:], currentValue, j, "right") // partTwoRight value
			ptv := partTwoVert(*inv, j, i, currentValue)           // partTwoVertical value

			result = append(result, ptl*ptr*ptv) // multiple each and store it in slice
		}
	}

	sort.Ints(result)
	fmt.Println("PART ONE:", *edges)
	fmt.Println("PART TWO:", result[len(result)-1])
}

func compareTopBottom(m [][]string, val int, i int, j int) int {
	result := 0
	left := m[j][:i] // m is an inverted matrix
	countLeft := compareLeftRight(left, val, i, "left")

	if countLeft == 0 {
		right := m[j][i+1:]
		countRight := compareLeftRight(right, val, i, "right")

		result += countRight
	}

	result += countLeft

	return result
}

func compareLeftRight(arr []string, val int, i int, position string) int {
	visibilityScore := 0

	if position == "left" {
		for j := i - 1; j >= 0; j-- {
			numeric, _ := strconv.Atoi(arr[j])
			if numeric >= val {
				visibilityScore = 0
				break
			} else {
				visibilityScore = 1
			}
		}
	}

	if position == "right" {
		for j := 0; j < len(arr); j++ {
			numeric, _ := strconv.Atoi(arr[j])
			if numeric >= val {
				visibilityScore = 0
				break
			} else {
				visibilityScore = 1
			}
		}
	}

	return visibilityScore
}

func invert(arr [][]string) [][]string {
	var result [][]string

	for i := 0; i < len(arr); i++ {
		var a []string
		for j := 0; j < len(arr[i]); j++ {
			a = append(a, arr[j][i])

		}
		result = append(result, a)
	}

	return result
}

func partTwo(arr []string, val int, i int, position string) int {
	visibleTrees := 0

	if position == "left" {
		for j := i - 1; j >= 0; j-- {
			numeric, _ := strconv.Atoi(arr[j])
			if numeric >= val {
				visibleTrees++
				break
			} else {
				visibleTrees++
			}
		}
	}

	if position == "right" {
		for j := 0; j < len(arr); j++ {
			numeric, _ := strconv.Atoi(arr[j])
			if numeric >= val {
				visibleTrees++
				break
			} else {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

func partTwoVert(m [][]string, j, i, val int) int {
	result := 1
	left := m[j][:i]
	countLeft := partTwo(left, val, i, "left")

	right := m[j][i+1:]
	countRight := partTwo(right, val, i, "right")

	result *= countLeft * countRight

	return result
}

func printJson(m interface{}) {
	json, err := json.MarshalIndent(m, "", " ")
	check(err)

	fmt.Println(string(json))
}
