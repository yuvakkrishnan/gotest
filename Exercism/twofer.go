// Instructions
// Two-fer or 2-fer is short for two for one. One for you and one for me.

// Given a name, return a string with the message:

// One for name, one for me.
// Where "name" is the given name.

// However, if the name is missing, return the string:

// One for you, one for me.
// Here are some examples:

// Name	String to return
// Alice	One for Alice, one for me.
// Bob	One for Bob, one for me.
// One for you, one for me.
// Zaphod	One for Zaphod, one for me.

package main

import "fmt"

func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}

func main() {

	result := ShareWith("Alice")
	fmt.Println(result)
}
