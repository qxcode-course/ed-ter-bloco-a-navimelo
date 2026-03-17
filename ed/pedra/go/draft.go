package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scan(&n)

    vencedor := -1
    diferenca:= 101

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

        if a >= 10 && b >= 10{
            dif:= int(math.Abs(float64(a - b)))

            if vencedor == -1 || dif < diferenca{
                diferenca = dif
                vencedor = i
            }
        }
	}

    if vencedor != -1{
        fmt.Println(vencedor)
    } else {
        fmt.Println("sem ganhador")
    }
}
