package main

import "fmt"

var c = make(map[int]int)

func coelhos(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 || n == 4 {
		return 2
	}

	if val, ok := c[n]; ok {
		return val
	}

	c[n] = coelhos(n-1) + coelhos(n-2) - coelhos(n-4)
	return c[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(coelhos(n))
}
