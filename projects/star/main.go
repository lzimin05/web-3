package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.Join(strings.Split(str, ""), "*")
	fmt.Println(str)
}
