package main
import "fmt"

func somaconj(elementos []int, idx int, soma int, alvo int) bool{
    if soma == alvo{
        return true
    }
    if soma > alvo || idx == len(elementos){
        return false
    }
    if somaconj(elementos, idx + 1, soma + elementos[idx], alvo){
        return true
    }
    if somaconj(elementos, idx + 1, soma, alvo){
        return true
    }
    return false
}

func main() {
    var n, k int

    if _, err := fmt.Scan(&n, &k); err != nil{
        return
    }

    elementos := make([]int, n)
    for i := 0; i < n; i++{
        fmt.Scan(&elementos[i])
    }

    if somaconj(elementos, 0, 0, k){
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}