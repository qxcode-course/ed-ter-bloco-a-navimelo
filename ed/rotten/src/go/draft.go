package main
import "fmt"

type posicao struct {
	linha int
	coluna int
}

func lpodre (grid [][]int) int{
	if len(grid) == 0 {
		return 0
	}
	l := len(grid)
	c := len(grid[0])
	
	fila := []posicao{}
	laranjaf := 0
	min := 0

	for i := 0; i < l; i++{
		for j := 0; j < c; j++ {
			switch grid[i][j] {
			case 2:
				fila = append(fila, posicao{i, j})
			case 1:
				laranjaf++
			}
		}
	}

	if laranjaf == 0 {
		return 0
	}

	direcao := [4]posicao{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(fila) > 0 && laranjaf > 0{
		t := len(fila)
		min++
		for i := 0; i < t; i++ {
			p := fila[0]
			fila = fila[1:]

			for _, d := range direcao {
				lin := p.linha + d.linha
				col := p.coluna + d.coluna

				if lin >= 0 && lin < l && col >= 0 && col < c && grid[lin][col] == 1 {
					grid[lin][col] = 2
					laranjaf--
					fila = append(fila, posicao{lin, col})
				}
			}
		}
	}
	if laranjaf == 0 {
		return min
	}
	return -1
}

func main() {
	var l, c int
	fmt.Scan(&l, &c)

	grid := make([][]int, l)
	for i := 0; i < l; i++ {
		grid[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	resultado := lpodre(grid)
	fmt.Println(resultado)
}
