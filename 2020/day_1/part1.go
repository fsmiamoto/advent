package main

import "fmt"

func main() {
	numbers := make(map[int]struct{})

	for {
		var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}

		numbers[n] = struct{}{}
	}

	for k := range numbers {
		_, ok := numbers[2020-k]
		if ok {
			fmt.Println(k * (2020 - k))
			break
		}

	}
}
