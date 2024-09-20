package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	var max rune
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	fmt.Printf("%c", max)
}
