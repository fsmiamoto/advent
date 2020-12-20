package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var bags map[string]map[string]int
var result map[string]bool

type Graph map[string][]string

func main() {
	file := readInputToString()

	rules := strings.Split(file, "\n")

	bags = make(map[string]map[string]int)

	graph := make(Graph)

	for _, r := range rules {
		color, children := getBags(r)
		for k := range children {
			graph[k] = append(graph[k], color)
		}
	}

	visited := make(map[string]struct{})
	fmt.Println(findThemAll(graph, "shiny gold", visited) - 1)
}

func findThemAll(g Graph, color string, visited map[string]struct{}) int {
	if _, ok := visited[color]; ok {
		return 0
	}

	visited[color] = struct{}{}

	total := 1
	for _, neighbor := range g[color] {
		total += findThemAll(g, neighbor, visited)
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
		count, _ := strconv.Atoi(s[0])
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
