package main

import (
	"fmt"
	"strings"
)

func main() {
	validPasswordCount := 0

	for {
		var min, max int
		var letter, password string
		_, err := fmt.Scanf("%d-%d %v %s", &min, &max, &letter, &password)
		if err != nil {
			break
		}

		letter = strings.TrimRight(letter, ":")

		count := 0
		for _, v := range strings.Split(password, "") {
			if v == letter {
				count += 1
			}
		}

		if count >= min && count <= max {
			validPasswordCount += 1
		}
	}

	fmt.Println(validPasswordCount)
}
