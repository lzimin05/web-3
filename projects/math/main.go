package main

import (
	"fmt"
	"math"
)

var k, p, v float64

// Начало решения задачи
func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}

// Конец решения

func main() {
	fmt.Scan(&k, &p, &v)
	fmt.Println(T())
}
