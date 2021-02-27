package main

import (
	"fmt"
)

func anagram (strings []string) {
	checkMapp := make(map[string]bool)

	iteration := 0

	for iteration < len(strings) {
		if checkMapp[strings[iteration]] { // check if already checked or already added
			iteration++
			continue
		}

		checkMapp[strings[iteration]] = true
		result := make([]string, 0)
		result = append(result, strings[iteration])

		for _, value := range strings {
			if checkMapp[value] { // check if it's already checked
				continue
			}

			if len(strings[iteration]) != len(value) { // check for len of words
				continue
			}

			totalIteration := 0
			totalValue := 0
			for i := 0; i < len(strings[iteration]); i++ {
				totalIteration += int(strings[iteration][i])
				totalValue += int(value[i])
			}
			if totalIteration != totalValue { // check total words
				continue
			}

			checkMapp[value] = true
			result = append(result, value)
		}
		fmt.Println(result)
		iteration++
	}
}

func main() {
	arr := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	anagram(arr)
}
