package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const length = 12 // hard coded

func main() {
	s := bufio.NewScanner(os.Stdin)

	ones := make([]uint, length)
	zeros := make([]uint, length)

	for s.Scan() {
		bits := strings.Split(s.Text(), "")
		for i, b := range bits {
			if b == "1" {
				ones[i]++
			} else {
				zeros[i]++
			}
		}
	}

	var gamma int
	for i := 0; i < length; i++ {
		if ones[i] > zeros[i] {
			gamma += (1 << (length - i - 1))
		}
	}

	epsilon := gamma ^ 0xFFF

	fmt.Println(gamma)
	fmt.Println(gamma * epsilon)
}
