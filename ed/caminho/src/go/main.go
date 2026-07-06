package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{l: p.l - 1, c: p.c}
		{l: p.l + 1, c: p.c}
		{l: p.l, c: p.c - 1}
		{l: p.l, c: p.c + 1}
	} 
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos){
	
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	if scanner.Scan(){
		line := scanner.Text()
		fmt.Sscanf(line, "%d %d", &nl, &nc)
	}

	mat := make([])
}				