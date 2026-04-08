package main

import "fmt"

func coelho(n, m int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return coelho(n-1, m) + (m * coelho(n-2, m))
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(coelho(n, m))
}
