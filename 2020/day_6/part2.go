package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file := readInputToString()

	groups := strings.Split(file, "\n\n")

	total := 0

	for i := range groups {
		groupAnswers := make(map[string]int)

		people := strings.Split(groups[i], "\n")
		for j := range people {
			answers := strings.Split(people[j], "")
			for k := range answers {
				groupAnswers[answers[k]] += 1
			}
		}

		for _, count := range groupAnswers {
			if count == len(people) {
				total += 1
			}
		}
	}

	fmt.Println(total)
}

func readInputToString() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}
