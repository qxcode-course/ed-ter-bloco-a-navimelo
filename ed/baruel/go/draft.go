package main
import "fmt"
func main() {
    var qtd_album, qtd_possui int
    fmt.Scan(&qtd_album, &qtd_possui)
    figuras := make ([]int, qtd_possui)

    for i := 0; i < qtd_possui; i++ {
        fmt.Scan(&figuras[i])
    }

}
