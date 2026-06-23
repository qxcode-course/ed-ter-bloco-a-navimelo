package main
import "fmt"

func dinheiro(igreja int, c int) float64 {
    if igreja == 0{
        return 0.0
    }
    return (dinheiro(igreja-1, c) + float64(c)) / 2.0
}

func main() {
    var n, c int
    fmt.Scan(&n, &c)
    
    doacao := dinheiro(n, c)
    fmt.Printf("%.2f\n", doacao)
}