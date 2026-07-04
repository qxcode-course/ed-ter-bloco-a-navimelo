package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct{
	l, c int
}

func resolver(grid [][]rune, nl, nc int, comeco, fim Pos){
	cam := NewStack[Pos]()
	beco := NewStack[Pos]()

	cam.Push(comeco)
	for !cam.IsEmpty(){
		atual := cam.Top()

		if atual.l == fim.l && atual.c == fim.c{
			grid[atual.l][atual.c] = '.'
			break
		}

		grid[atual.l][atual.c] = '.'

		dirl := []int{-1, 1, 0, 0}
		dirc := []int{0, 0, 1, -1}

		var viz []Pos

		for i := 0; i < 4; i++{
			proxl := atual.l + dirl[i]
			proxc := atual.c + dirc[i]

			if proxl >= 0 && proxl < nl && proxc >= 0 && proxc < nc{
				char := grid[proxl][proxc]

				if char != '#' && char != '.' && char != 'b'{
					viz = append(viz, Pos{proxl, proxc})
				}
			}
		}
		if len(viz) > 0{
			cam.Push(viz[0])
		} else{
			grid[atual.l][atual.c] = 'b'
			beco.Push(atual)
			cam.Pop()
		}
	}

	for !beco.IsEmpty(){
		b := beco.Pop()
		grid[b.l][b.c] = ' '
	}
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan(){
		return
	}
	var nl, nc int
	fmt.Sscan(scanner.Text(), &nl, &nc)

	grid := make([][]rune, 0, nl)
	for i := 0; i < nl; i++{
		if !scanner.Scan(){
		break
		}
		grid = append(grid, []rune(scanner.Text()))
	}

	var comeco, fim Pos

	for l := 0; l < nl; l++{
		for c := 0; c < nc; c++{
			if grid[l][c] == 'I'{
				comeco = Pos{l, c}
			} else if grid[l][c] == 'F'{
				fim = Pos{l, c}
			}
		}
	}

	resolver(grid, nl, nc, comeco, fim)
	for _, lin := range grid{
		fmt.Println(string(lin))
	}
}