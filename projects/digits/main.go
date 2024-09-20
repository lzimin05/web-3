package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	var maxNumber int
	maxNumber = 0
	fmt.Scan(&str)
	for _, symbol := range str {
		i, err := strconv.Atoi(string(symbol))
		if err != nil {
			panic(err)
		}
		maxNumber = max(maxNumber, i)
	}
	fmt.Println(maxNumber)
}
