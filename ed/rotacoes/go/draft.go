package main

import "fmt"

func main() {
	var t, r int
	fmt.Scan(&t, &r)

	vetor := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scan(&vetor[i])
	}

	rotacao := len(vetor)
	if rotacao > 0 {
		rotacao = r % len(vetor)
	}

	resultado := make([]int, len(vetor))
	for i := 0; i < len(vetor); i++ {
		novap := (i + rotacao) % len(vetor)
		resultado[novap] = vetor[i]
	}

	fmt.Printf("[ ")

	for i, valor := range resultado {
		fmt.Print(valor)
		if i < len(resultado)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println(" ]")
}
