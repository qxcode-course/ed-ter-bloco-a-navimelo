package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)

	x := (a + b + c) / 2

	area := math.Sqrt(x * (x - a) * (x - b) * (x - c))

	fmt.Printf("%.2f\n", area)
}
