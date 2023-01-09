package main

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.Replace(strings.ToLower(word), "-", "", -1)
	word = strings.Replace(strings.ToLower(word), " ", "", -1)

	letterFrequencies := make(map[rune]int)

	for _, letter := range word {
		letterFrequencies[letter]++
	}

	for _, frequency := range letterFrequencies {
		if frequency > 1 {
			return false
		}
	}

	return true
}

func main() {
	words := []string{"lumberjacks", "background", "downstream", "six-year-old", "isograms"}
	for _, word := range words {
		if IsIsogram(word) {
			fmt.Println(word, "is an isogram")
		} else {
			fmt.Println(word, "is not an isogram")
		}
	}
}
