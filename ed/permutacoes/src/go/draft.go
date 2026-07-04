package main
import (
    "fmt"
)

func permutar(a string, r string){
    if len(r) == 0{
        fmt.Println(a)
        return 
    }

    for i := 0; i < len(r); i++{
        let := string(r[i])
        resto := r[:i] + r[i + 1:]
        permutar(a + let, resto)
    }

}

func main(){
    var s string

    if _, err := fmt.Scan(&s); err != nil{
        return
    }

    lets := []rune(s)
    for i := 0; i < len(lets); i++{
        for j := i +1; j < len(lets); j++{
            if lets[i] > lets[j]{
                lets[i], lets[j] = lets[j], lets[i]
            }
        }
    }

    permutar("", string(lets))
}