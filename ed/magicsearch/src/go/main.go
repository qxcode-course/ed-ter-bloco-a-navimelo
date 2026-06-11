package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	if len(slice) == 0 {
		return 0
	}

	esquerda, direita := 0, len(slice)-1
	ultimo := -1
	for esquerda <= direita {
		meio := (esquerda + direita) / 2
		if slice[meio] == value {
			ultimo = meio
			esquerda = meio + 1
		} else if slice[meio] < value {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}

	if ultimo != -1 {
		return ultimo
	}

	return esquerda
	//_, _ = slice, value
	//return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
