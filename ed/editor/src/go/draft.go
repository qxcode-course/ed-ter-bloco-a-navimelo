package main
import (
    "bufio"
    "fmt"
    "os"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan(){
        leitor := scanner.Text(){
            var esq []rune
            var dir []rune

            for _, char := range leitor{
                if(char >= 'a' && char <= 'z') || char == '-'{
                    esq = append(esq, char)
                } else {
                    switch char{
                    case 'R':
                    case 'B':
                    case 'D':
                    case '>':
                    case '<':
                    }
                }
            }
        }
    }
}