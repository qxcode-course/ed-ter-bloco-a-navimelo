package main
import "fmt"
import "math"

type rainhas struct {
    n int
    m int
    c []int
}

func novas_rainhas(n int) *rainhas {
    return &rainhas{n: n, c: make([]int, n)}
}

func (r *rainhas) pode_colocar(l, c int) bool {
    for i := 0; i < l; i++ {
        if r.c[i] == c{
            return false
        }
        if math.Abs(float64(r.c[i]-c)) == math.Abs(float64(i-l)) {
            return false
        }
    }
    return true
}

func (r *rainhas) resolver(l int) {
    if l == r.n {
        r.m++
        return
    }
    for c := 0; c < r.n; c++{
        if r.pode_colocar(l, c){
            r.c[l] = c
            r.resolver(l + 1)
        }
    }
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }
    
    r := novas_rainhas(n)
    r.resolver(0)
    fmt.Println(r.m)
}
