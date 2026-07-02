package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func precedencia(pol string) int{
    switch pol{
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    case "^":
        return 3
    }
    return 0
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan(){
        exp := scanner.Text()
        elem := strings.Fields(exp)

        var resul []string
        var pilha []string

        for _, e := range elem{
            if e == "+" || e == "-" || e == "*" || e == "/" || e == "^"{
                for len(pilha) > 0 && precedencia(pilha[len(pilha)-1]) >= precedencia(e){
                    t := pilha[len(pilha) -1]
                    pilha = pilha[:len(pilha) -1]
                    resul = append(resul, t)
                }
                pilha = append(pilha, e)
            } else {
                resul = append(resul, e)
            }
        }

        for len(pilha) > 0{
            t := pilha[len(pilha)-1]
            pilha = pilha[:len(pilha)-1]
            resul = append(resul, t)
        }

        fmt.Println(strings.Join(resul, " "))
    }

}