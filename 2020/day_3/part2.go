package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file := readFileToString(filename)
	fmt.Println(
		countTrees(file, 0, 1, 1) *
			countTrees(file, 0, 3, 1) *
			countTrees(file, 0, 5, 1) *
			countTrees(file, 0, 7, 1) *
			countTrees(file, 0, 1, 2))
}

func countTrees(line string, index int, right int, down int) int {
	max := strings.Index(line, "\n")
	treeCount := 0

	if max == -1 {
		return treeCount
	}

	if line[index%max] == '#' {
		treeCount += 1
	}

	if down*(max+1) > len(line) {
		return treeCount
	}

	return treeCount + countTrees(line[down*(max+1):], index+right, right, down)
}

func readFileToString(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
