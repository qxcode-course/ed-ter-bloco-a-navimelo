package main
import "fmt"

func quadrado (n int) int {
    if n == 1 {
        fmt.Printf("%d^2 = %d\n", n, 1)
        return 1
    }
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)

    rec := quadrado (n-1)

    resultado := rec + 2*(n-1) + 1

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, n-1, n-1, resultado)

    return resultado
}

func main() {
    var x int 
    fmt.Scanln(&x)
    
    quadrado(x)
}