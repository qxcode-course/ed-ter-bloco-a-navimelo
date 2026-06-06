package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	vencedor := -1
	menorDiferenca := 101

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		if a >= 10 && b >= 10 {
			dif := a - b
			if dif < 0 {
				dif = -dif
			}

			if vencedor == -1 || dif < menorDiferenca {
				menorDiferenca = dif
				vencedor = i
			}
		}
	}

	if vencedor != -1 {
		fmt.Println(vencedor)
	} else {
		fmt.Println("sem ganhador")
	}
}