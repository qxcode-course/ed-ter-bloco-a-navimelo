package main
import (
    "fmt"
)

func detona(bomba [][]int) int{
    n := len(bomba)
    maxbomba := 0

    for i := 0; i < n; i++{
        vis := make([]bool, n)
        cont := dfs(i, bomba, vis)

        if cont > maxbomba{
            maxbomba = cont
        }
    }
    return maxbomba
}
func dfs(atual int, bomba[][]int, vis []bool) int{
    vis[atual] = true
    cont := 1

    a, b, r := bomba[atual][0], bomba[atual][1], bomba[atual][2]

    for prox := 0; prox < len(bomba); prox++{
        if vis[prox]{
            continue
        }
        x, y := bomba[prox][0], bomba[prox][1]

        dx := int64(a - x)
        dy := int64(b - y)
        distancia := (dx * dx) + (dy * dy)
        raio := int64(r) * int64(r)

        if distancia <= raio{
            cont += dfs(prox, bomba, vis)
        }
    }
    return cont
}

func main(){
    var n, c int
    fmt.Scan(&n, &c)

    bomba := make([][]int, n)
    for i := 0; i < n; i++{
        bomba[i] = make([]int, 3)
        fmt.Scan(&bomba[i][0], &bomba[i][1], &bomba[i][2])
    }
    fmt.Println(detona(bomba))
}