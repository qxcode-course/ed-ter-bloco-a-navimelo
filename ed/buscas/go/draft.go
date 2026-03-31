package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)

	entrada := make(map[string]int)

    for i := 0; i < n; i++ {
        var palavra string
        fmt.Scan(&palavra)
        entrada[palavra]++
    }

    fmt.Scan(&m)
    consultas := make([]string, m)

    for i := 0; i < m; i++ {
        fmt.Scan(&consultas[i])
    }

    for i, p := range consultas {
        fmt.Print(entrada[p])
        
        if i < m-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println()

}
