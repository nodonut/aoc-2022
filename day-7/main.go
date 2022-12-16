package main

import (
	"bufio"
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

type child struct {
	Name string
	Size int
}

type directory struct {
	Name           string
	Size           int
	Parent         *directory
	Subdirectories map[string]*directory
	Files          map[string]*child
}

func (d *directory) addFile(newFile *child) {
	if d.Files != nil {
		d.Files[newFile.Name] = newFile
	} else {
		d.Files = make(map[string]*child)
		d.Files[newFile.Name] = newFile
	}
}

func (d *directory) addSubdirectory(sub *directory) {
	if d.Subdirectories != nil {
		d.Subdirectories[sub.Name] = sub
	} else {
		d.Subdirectories = make(map[string]*directory)
		d.Subdirectories[sub.Name] = sub
	}
}

func (d *directory) getDirectorySize() int {
	sum := 0
	for _, file := range d.Files {
		sum += file.Size
	}
	for _, subdirectory := range d.Subdirectories {
		sum += subdirectory.getDirectorySize()
	}
	d.Size = sum

	return sum
}

func main() {
	file, err := os.Open("prod.txt")
	check(err)
	defer file.Close()

	scnr := bufio.NewScanner(file)
	root := &directory{Name: "/"}
	currentDirectory := root

	for scnr.Scan() {
		line := scnr.Text()

		if isCd(line) {
			if isCdRoot(line) {
				currentDirectory = root
			} else if isCdPrev(line) {
				currentDirectory = currentDirectory.Parent
			} else {
				subName := strings.Split(line, " ")[2]
				for _, sub := range currentDirectory.Subdirectories {
					if sub.Name == subName {
						currentDirectory = sub
					}
				}
			}
		}

		if isDir(line) {
			dir := &directory{Name: strings.Split(line, " ")[1], Parent: currentDirectory}
			currentDirectory.addSubdirectory(dir)
		}

		if size, _ := strconv.Atoi(strings.Split(line, " ")[0]); size != 0 {
			f := strings.Split(line, " ")[1]
			file := &child{Name: f}
			file.Size = size
			currentDirectory.addFile(file)
		}

	}

	root.getDirectorySize()

	var dirSizes = make(map[string]int)
	for s, sub := range root.Subdirectories {
		dirSizes["/"+s] = sub.Size
		getSubdirectorySizes(sub, dirSizes, s)
	}

	part1(dirSizes)
	part2(root, dirSizes)
}

func part1(dirSizes map[string]int) {
	result := 0
	for _, v := range dirSizes {
		if v <= 100000 {
			result += v
		}
	}

	fmt.Println("PART 1: ", result)
}

func part2(root *directory, dirSizes map[string]int) {
	maxSize := 70000000
	unusedSpaceNeeded := 30000000
	unusedSpace := maxSize - root.Size

	result := []int{}
	for _, v := range dirSizes {
		if unusedSpace+v >= unusedSpaceNeeded {
			result = append(result, v)
		}
	}

	sort.Ints(result)

	fmt.Println("PART 1: ", result[0])
}

func isCd(line string) bool {
	return strings.Split(line, " ")[1] == "cd"
}

func isCdRoot(line string) bool {
	return isCd(line) && strings.Split(line, " ")[2] == "/"
}

func isCdPrev(line string) bool {
	return isCd(line) && strings.Split(line, " ")[2] == ".."
}

func isDir(line string) bool {
	return strings.Split(line, " ")[0] == "dir"
}

func getSubdirectorySizes(d *directory, result map[string]int, parent string) {
	for s, subdirectory := range d.Subdirectories {
		result[parent+"/"+s] = subdirectory.Size
		getSubdirectorySizes(subdirectory, result, s)
	}
}
