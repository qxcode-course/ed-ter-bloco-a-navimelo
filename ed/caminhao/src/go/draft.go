package main
import "fmt"
func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil{
        return
    }

    total := 0
    tanque := 0
    comeco := 0
    for i := 0; i < n; i++{
        var gasolina, distancia int
        fmt.Scan(&gasolina, &distancia)

        saldo := gasolina - distancia
        total += saldo
        tanque += saldo

        if tanque < 0{
            comeco = i + 1
            tanque = 0
        }
    }

    fmt.Println(comeco)
}