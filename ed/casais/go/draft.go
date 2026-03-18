package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    var a [50]int
    var f = 0
    var m = 0

    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
        if a[i] < 0 {
            f++
        } else{
            m++
        }
    }

    var casais = 0

    if f > m{
        casais = m
    } else if f < m {
        casais = f 
    } else if f == m {
        casais = f
    }

    fmt.Println(casais)
    
}
