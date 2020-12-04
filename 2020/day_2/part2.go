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
		letters := strings.Split(password, "")

		if letters[min-1] == letter {
			count += 1
		}
		if letters[max-1] == letter {
			count += 1
		}

		if count == 1 {
			validPasswordCount += 1
		}
	}

	fmt.Println(validPasswordCount)
}
