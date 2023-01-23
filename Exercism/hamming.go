// Instructions
// Calculate the Hamming Distance between two DNA strands.

// Your body is made up of cells that contain DNA. Those cells regularly wear out and need replacing, which they achieve by dividing into daughter cells. In fact, the average human body experiences about 10 quadrillion cell divisions in a lifetime!

// When cells divide, their DNA replicates too. Sometimes during this process mistakes happen and single pieces of DNA get encoded with the incorrect information. If we compare two strands of DNA and count the differences between them we can see how many mistakes occurred. This is known as the "Hamming Distance".

// We read DNA using the letters C,A,G and T. Two strands might look like this:

// GAGCCTACTAACGGGAT
// CATCGTAATGACGGCCT
// ^ ^ ^  ^ ^    ^^
// They have 7 differences, and therefore the Hamming Distance is 7.

// The Hamming Distance is useful for lots of things in science, not just biology, so it's a nice phrase to be familiar with :)

// Implementation notes
// The Hamming distance is only defined for sequences of equal length, so an attempt to calculate it between sequences of different lengths should not work.

// You may be wondering about the cases_test.go file. We explain it in the leap exercise.

package main

import "fmt"

func Distance(i, j string) (int, error) {

	if len(i) != len(j) {
		return 0, fmt.Errorf("The Distance is not equal", len(i), "!=", len(j))
	}
	var distance int
	for v := range i {
		if i[v] != j[v] {
			distance++
		}
	}
	return distance, nil
}
func main() {
	var wordOne string
	var wordTwo string
	fmt.Println("Enter 1st word to hamp ")
	fmt.Scanln(&wordOne)
	fmt.Println("Enter 2nd word to hamp ")
	fmt.Scanln(&wordTwo)

	dist, err := Distance(wordOne, wordTwo)
	if err != nil {
		panic(err)
	}
	fmt.Println(dist)

}

// Hamming with Go Routine

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func hammingWords(word1, word2 string) (int, error) {

// 	if len(word1) != len(word2) {
// 		return 0, fmt.Errorf("Error len %d != len %d", len(word1), len(word2))
// 	}
// 	distance := 0
// 	for v := range word1 {
// 		if word1[v] != word2[v] {
// 			distance++
// 		}

// 	}
// 	return distance, nil
// }

// func main() {
// 	var mu sync.WaitGroup
// 	mu.Add(1)
// 	go func() {
// 		defer mu.Done()
// 		callstat, err := hammingWords("hello", "Hello")
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(callstat)
// 	}()
// 	mu.Wait()

// }
