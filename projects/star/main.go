package main

import (
	"fmt"
	"unicode/utf8" // 1 вариант
	//"strings" // 2 вариант
)

func main() {
	var str, res string
	fmt.Scan(&str)
	for index, symbol := range str { // 1 вариант с циклом
		res += string(symbol)
		if index != utf8.RuneCountInString(str)-1 {
			res += string('*')
		}
	}
	//res = strings.Join(strings.Split(str, ""), string('*')) // 2 вариант
	fmt.Println(res)
}
