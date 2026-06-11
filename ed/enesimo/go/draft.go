package main

import "fmt"

func eh_primo(x int, div int) bool {

	if x < 2 {
		return false
	}
	if div == x {
		return true
	}
	if x%div == 0 {
		return false
	}

	return eh_primo(x, div+1)
}

func proximo (x int, tl int, cont int) int {
	if eh_primo(tl , 2) {
		cont ++
	}
	if cont == x {
		return tl
	}
	return proximo (x, tl + 1, cont)
}

func enesimo (x int) int {
	return proximo(x, 2, 0)
}

func main() {
	var x int
	fmt.Scan(&x)
	
	fmt.Println(enesimo(x))
}
