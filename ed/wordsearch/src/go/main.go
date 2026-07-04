package main

import (
	"bufio"
	"fmt"
	"os"
)

func busca(grid [][]byte, l, c int, palavra string, idx int) bool{
	if idx == len(palavra){
		return true
	}
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) || grid[l][c] != palavra[idx]{
		return false
	}
	let := grid[l][c]
	grid[l][c] = ' '

	dirl := []int{-1, 1, 0, 0}
	dirc := []int{ 0, 0, 1, -1}

	for i := 0; i <4; i++{
		proxl := l + dirl[i]
		proxc := c + dirc[i]

		if busca(grid, proxl, proxc, palavra, idx + 1){
			return true
		}
	}
	grid[l][c] = let
	return false
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 || len(word) == 0{
		return false
	}
	for l := 0; l < len(grid); l++{
		for c := 0; c < len(grid[0]); c++{
			if grid[l][c] == word[0]{
				if busca(grid, l, c, word, 0){
					return true
				}
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
