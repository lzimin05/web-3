package main

import (
	"fmt"
	"unicode"
)

func SqrtDigit(a string) {
	for _, i := range a {
		if unicode.IsDigit(i) {
			c := int(i - 48)
			fmt.Print(c * c)
		}
	}
}

func main() {
	var a string
	fmt.Scan(&a)
	SqrtDigit(a)
}
