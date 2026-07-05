package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0{
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])

	mem := make([][]int, m)

	for i := range mem{
		mem[i]  = make([]int, n)
	}
	direcao := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}

	var dfs func(r, c int) int
	dfs = func(r, c int) int{
		if mem[r][c] != 0 {
			return mem[r][c]
		}
		tam := 1
		for _, d := range direcao{
			nr, nc := r + d[0], c + d[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n && matrix[nr][nc] > matrix[r][c]{
				caminho := 1 + dfs(nr, nc)
				if caminho > tam{
					tam = caminho
				}
			}
		}
		mem[r][c] = tam
		return tam
	}
	resultado := 0

	for i := 0; i < m; i ++{
		for j := 0; j < n; j++{
			camin := dfs(i, j)
			if camin > resultado{
				resultado = camin
			}
		}
	}
	return resultado
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
