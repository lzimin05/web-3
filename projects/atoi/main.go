package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str, res string
	fmt.Scan(&str)
	for _, symbol := range str {
		i, err := strconv.Atoi(string(symbol))
		if err != nil {
			panic(err)
		}
		res += strconv.Itoa(i * i)
	}
	fmt.Println(res)
}
