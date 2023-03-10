package main

import (
	"fmt"
	"strings"
)

type IsogramChecker struct{}

func (i *IsogramChecker) Check(word string) bool {
	word = strings.Replace(strings.ToLower(word), "-", "", -1)
	word = strings.Replace(strings.ToLower(word), " ", "", -1)

	//make(map[rune]int) is used to create a new empty map with keys of type rune and values of type int.
	letterFrequencies := make(map[rune]int)

	for _, letter := range word {
		//On each iteration, letterFrequencies[letter]++ increments the value associated with the letter key in the map by 1. If the key does not exist in the map yet, it will be created with the initial value of 1.
		letterFrequencies[letter]++
	}
		if frequency > 1 {
			return false
		}
	}

	return true
}

func main() {
	ic := &IsogramChecker{}
	words := []string{"lumberjacks", "background", "downstream", "six-year-old", "isograms"}
	for _, word := range words {
		if ic.Check(word) {
			fmt.Println(word, "is an isogram")
		} else {
			fmt.Println(word, "is not an isogram")
		}
	}
}

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func IsIsogram(word string) bool {
// 	word = strings.Replace(strings.ToLower(word), "-", "", -1)
// 	word = strings.Replace(strings.ToLower(word), " ", "", -1)

// 	letterFrequencies := make(map[rune]int)

// 	for _, letter := range word {
// 		letterFrequencies[letter]++
// 	}

// 	for _, frequency := range letterFrequencies {
// 		if frequency > 1 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func main() {
// 	words := []string{"lumberjacks", "background", "downstream", "six-year-old", "isograms"}
// 	for _, word := range words {
// 		if IsIsogram(word) {
// 			fmt.Println(word, "is an isogram")
// 		} else {
// 			fmt.Println(word, "is not an isogram")
// 		}
// 	}
// }
