package main

import "fmt"

func imprimir(lista []int) {
	if len(lista) == 0 {
		fmt.Println("N")
		return
	}
	for i, v := range lista {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	var qtdAlbum, qtdPossui int
	if _, err := fmt.Scan(&qtdAlbum, &qtdPossui); err != nil {
		return
	}

	album := make([]bool, qtdAlbum+1)
	var repetidos []int

	for i := 0; i < qtdPossui; i++ {
		var figura int
		fmt.Scan(&figura)

		if album[figura] {
			repetidos = append(repetidos, figura)
		} else {
			album[figura] = true
		}
	}

	var faltantes []int
	for i := 1; i <= qtdAlbum; i++ {
		if !album[i] {
			faltantes = append(faltantes, i)
		}
	}

	imprimir(repetidos)
	imprimir(faltantes)
}
