package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0{
		return 0
	}
	total := 0

	var busca func(r, c int)
	busca = func(r, c int){
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == '0'{
			return
		}
		grid[r][c] = '0'

		busca(r + 1, c)
		busca(r - 1, c)
		busca(r, c + 1)
		busca(r, c - 1)
	}
	for i := 0; i < len(grid);  i++{
		for j := 0; j < len(grid[0]);  j++{
			if grid[i][j] == '1'{
				total++
				busca(i, j)
			}
		}
	}

	return total
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
