package main

import "fmt"

func resolver(matriz [][]rune, index int) bool {
    n := len(matriz)
    if index == n*n {
        return true
    }
    lin, col := index/n, index%n
    if matriz[lin][col] != '.' {
        return resolver(matriz, index+1)
    }
    for num := '1'; num <= rune('0'+n); num++ {
        if valido(matriz, lin, col, num) {
            matriz[lin][col] = num
            if resolver(matriz, index+1) {
                return true
            }
            matriz[lin][col] = '.'
        }
    }
    return false
}

func valido(matriz [][]rune, lin, col int, num rune) bool {
    n := len(matriz)

    for c := 0; c < n; c++ {
        if matriz[lin][c] == num {
            return false
        }
    }

    for l := 0; l < n; l++ {
        if matriz[l][col] == num {
            return false
        }
    }

    tam := 2
    if n == 9 {
        tam = 3
    }
    linInicio := (lin / tam) * tam
    colInicio := (col / tam) * tam
    for i := 0; i < tam; i++ {
        for j := 0; j < tam; j++ {
            if matriz[linInicio+i][colInicio+j] == num {
                return false
            }
        }
    }
    return true
}

func main() {
    var n int
    fmt.Scan(&n)

    matriz := make([][]rune, n)
    for i := 0; i < n; i++ {
        var linha string
        fmt.Scan(&linha)
        matriz[i] = []rune(linha)
    }

    if resolver(matriz, 0) {
        for i := 0; i < n; i++ {
            fmt.Println(string(matriz[i]))
        }
    }
}