package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	if m == 0 || n == 0{
		return
	}
	var busca func(r, c int)
	busca = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != 'O'{
			return 
		}
		board[r][c] = 'V'
		busca(r + 1, c)
		busca(r - 1, c)
		busca(r, c + 1)
		busca(r, c - 1)
	}
	for i:= 0; i < m; i++{
		busca(i, 0)
		busca(i, n - 1)
	}
	for j:= 0; j < n; j++{
		busca(0, j)
		busca(m - 1, j)
	}
	for i := 0; i < m; i++{
		for j:= 0; j < n; j++{
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			} else if board[i][j] == 'V'{
				board[i][j] = 'O'
			}
		}
	}
}
// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
