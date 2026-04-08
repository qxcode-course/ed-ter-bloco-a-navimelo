package main
import "fmt"

func binomial(n, k int) int {
    if k == 0 || k == n {
        return 1
    }
    return binomial(n-1, k-1) + binomial(n-1, k)
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    fmt.Println(binomial(n, k))
}
