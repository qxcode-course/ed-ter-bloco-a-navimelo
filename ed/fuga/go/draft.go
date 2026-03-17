package main

import "fmt"
func main() {
    var h, p, f, d int

    fmt.Scan(&h, &p, &f, &d)

    if d == -1{
        if (f > h && f > p && p >= h) || (f < h && (p >= h || p < f)){
            fmt.Println("N")
        } else {
            fmt.Println("S")
        }
    } else {
        if (f < h && p > f && p <= h) || (f > h && (p<= h || p > f)){
             fmt.Println("N")
        } else {
            fmt.Println("S")
        }
    }
}
