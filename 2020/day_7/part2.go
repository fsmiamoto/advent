package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var bags map[string]map[string]int
var result map[string]bool

type Graph map[string]map[string]int

func main() {
	file := readInputToString()

	rules := strings.Split(file, "\n")

	bags = make(map[string]map[string]int)

	graph := make(Graph)

	for _, r := range rules {
		parentColor, childrenBags := getBags(r)
		graph[parentColor] = childrenBags
	}

	fmt.Println(findThemAll(graph, "shiny gold") - 1)
}

func findThemAll(g Graph, color string) int {
	total := 1
	for neighbor, count := range g[color] {
		total += count * findThemAll(g, neighbor)
	}

	return total
}

func getBags(s string) (string, map[string]int) {
	if s == "" {
		return "", nil
	}

	if strings.Contains(s, "no other bags") {
		return "", nil
	}

	trimmed := strings.TrimRight(s, ".")
	splitted := strings.Split(trimmed, " bags contain ")

	baseBag := splitted[0]
	result := make(map[string]int)

	for _, color := range strings.Split(splitted[1], ", ") {
		s := strings.SplitAfterN(color, " ", 2)
		count, _ := strconv.Atoi(strings.Trim(s[0], " "))
		s = strings.Split(s[1], " ")
		result[s[0]+" "+s[1]] += count
	}

	return baseBag, result
}

func getBagColor(s string) string {
	splitted := strings.Split(s, " ")
	if len(splitted) != 3 {
		return strings.Join(splitted, " ")
	}
	return strings.Join(strings.Split(s, " ")[:2], " ")
}

func readInputToString() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}
