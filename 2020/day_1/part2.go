package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers []int

	for {
		var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}

		numbers = append(numbers, n)
	}

	sort.Ints(numbers)

	var l, r, sum int

	for i := 0; i < len(numbers)-2; i++ {
		l = i + 1
		r = len(numbers) - 1

		for l < r {
			sum = numbers[i] + numbers[l] + numbers[r]
			if sum == 2020 {
				product := numbers[i] * numbers[l] * numbers[r]
				fmt.Printf("%d * %d * %d = %d \n", numbers[i], numbers[l], numbers[r], product)
				break
			}

			if sum < 2020 {
				l += 1
			} else {
				r -= 1
			}
		}
	}
}
