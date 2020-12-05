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
	fmt.Println(countTrees(file, 0))
}

const slope = 3

func countTrees(line string, index int) int {
	max := strings.Index(line, "\n")
	treeCount := 0

	if max == -1 {
		return treeCount
	}

	if line[index%max] == '#' {
		treeCount += 1
	}

	return treeCount + countTrees(line[max+1:], index+slope)
}

func readFileToString(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
