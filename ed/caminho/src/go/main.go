package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{l: p.l - 1, c: p.c},
		{l: p.l + 1, c: p.c},
		{l: p.l, c: p.c - 1},
		{l: p.l, c: p.c + 1},
	} 
}

func inside(grid [][]rune, pos Pos) bool {
	//nrows := len(grid)
	//ncols := len(grid[0])
	//return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
	return pos.l >= 0 && pos.l < len(grid) && pos.c >= 0 && pos.c < len(grid[pos.l])
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos){
	queue := NewQueue[Pos]()
	queue.Enqueue(startPos)

	vis := make(map[Pos]bool)
	vis[startPos] = true

	ant := make(map[Pos]Pos)

	final := false

	for !queue.IsEmpty(){
		atual, _ := queue.Dequeue()

		if atual == endPos{
			final = true
			break
		}
		for _, viz := range atual.getNeig(){
			if inside(grid, viz) && !match(grid, viz, '#') && !vis[viz]{
				vis[viz] = true
				ant[viz] = atual
				queue.Enqueue(viz)
			}
		}
	}
	if final{
		atual := endPos
		for {
			grid[atual.l][atual.c] = '.'
			if atual == startPos {
				break
			}
			atual = ant[atual]
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	if scanner.Scan(){
		line := scanner.Text()
		line = strings.TrimRight(line, "\r\n")
		fmt.Sscanf(line, "%d %d", &nl, &nc)
	}

	mat := make([][] rune, nl)

	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		line = strings.TrimRight(line, "\r\n")
		
		runes := []rune(line)
		for len(runes) < nc{
			runes = append(runes, ' ')
		}
		mat[i] = runes
	}
	var ini, fim Pos

	for l := range nl{
		for c := range nc {
			if mat[l][c] == 'I'{
				ini = Pos{l,c}
			}
			if mat[l][c] == 'F'{
				fim = Pos{l,c}
			}
		}
	}
	search(mat, ini, fim)

	for _, line := range mat{
		fmt.Println(string(line))
	}
}				