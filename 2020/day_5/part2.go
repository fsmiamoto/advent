package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file := readInputToString()
	seats := strings.Split(file, "\n")

	var ids []int
	for i := range seats {
		row, col := getRowAndColFromSeat(seats[i])
		id := getSeatID(row, col)
		if id == 0 {
			continue
		}
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1]-ids[i] > 1 {
			fmt.Println(ids[i] + 1)
			break
		}
	}
}

func getSeatID(row, col int) int {
	return row*8 + col
}

func getRowAndColFromSeat(s string) (int, int) {
	if len(s) == 0 {
		return 0, 0
	}

	rowChars := s[:7]
	replacer := strings.NewReplacer("F", "0", "B", "1")
	rowChars = replacer.Replace(rowChars)
	row, _ := strconv.ParseInt(rowChars, 2, 64)

	colChars := s[7:]
	replacer = strings.NewReplacer("R", "1", "L", "0")
	colChars = replacer.Replace(colChars)
	col, _ := strconv.ParseInt(colChars, 2, 64)

	return int(row), int(col)
}

func readInputToString() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}
