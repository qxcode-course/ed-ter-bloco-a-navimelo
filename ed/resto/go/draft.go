package main
import "fmt"

func divisao(num int){
    if num == 0{
        return
    }
    q := num/2
    r := num % 2

    divisao(q)

    fmt.Println(q, r)
}

func main() {
    var i int
    fmt.Scan(&i)

    divisao(i)
}
