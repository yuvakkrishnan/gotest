package main

import "fmt"

func distance(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return 0, fmt.Errorf("strings have different length: %d != %d", len(s1), len(s2))
	}
	var distance int
	for i := range s1 {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance, nil
}

func main() {
	s1 := "def"
	s2 := "def"
	distance, err := distance(s1, s2)
	if err != nil {
		fmt.Printf("Error calculating distance: %v\n", err)
		return
	}
	fmt.Printf("Distance between %q and %q: %d\n", s1, s2, distance)

}
