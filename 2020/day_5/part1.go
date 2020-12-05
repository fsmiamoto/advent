package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file := readInputToString()
	seats := strings.Split(file, "\n")
	max := 0
	for i := range seats {
		row, col := getRowAndColFromSeat(seats[i])
		id := getSeatID(row, col)
		fmt.Println(id)
		if id > max {
			max = id
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
