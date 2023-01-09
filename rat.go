package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	var y int

	y = x * *&x
	fmt.Println(y)

}
