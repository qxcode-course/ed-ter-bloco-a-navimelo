package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	nl := len(grid)
	if nl == 0{
		return
	}
	nc := len(grid[0])
	
	for !stack.IsEmpty(){
		atual := stack.Pop()

		if atual.l < 0 || atual.l >= nl || atual.c < 0 || atual.c >= nc {
			continue
		}

		if grid[atual.l][atual.c] == '#'{
			grid[atual.l][atual.c] = 'o'

			stack.Push(Pos{atual.l - 1, atual.c})
			stack.Push(Pos{atual.l + 1, atual.c})
			stack.Push(Pos{atual.l, atual.c - 1})
			stack.Push(Pos{atual.l, atual.c + 1})
		}
	}

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}