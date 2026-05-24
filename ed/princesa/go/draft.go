package main

import "fmt"

func print_jog(j []int, e int){
    fmt.Print("[ ")
    for i, elem := range j{
        if elem == 0{
            continue
        }
        fmt.Print(elem)
        if i == e {
            fmt.Print(">")
        }
        fmt.Print(" ")
    }
    fmt.Print("]\n")
}

func procurar_vivo(j []int, e int) int{
    for {
       e = (e + 1) % len(j)
        if j [e] != 0 {
            return e
        }
    }
}

func main() {
	var q, e int
	fmt.Scan(&q, &e)

	j := make([]int, 0, q)
    
	for i := 1; i <= q; i++ {
		j = append(j, i)
	}
    e -= 1
    for range q - 1{
        print_jog(j, e)
        vai_morrer := procurar_vivo(j, e)
        j[vai_morrer] = 0
        e = procurar_vivo(j, e)
    }
    print_jog(j, e)
}
