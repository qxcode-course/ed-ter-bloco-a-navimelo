package main
import (
    "bufio"
    "fmt"
    "os"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan(){
        leitor := scanner.Text()
            var esq []rune
            var dir []rune

            for _, char := range leitor{
                if(char >= 'a' && char <= 'z') || char == '-'{
                    esq = append(esq, char)
                } else {
                    switch char{
                    case 'R':
                        esq = append(esq, '\n')
                    case 'B':
                        if len(esq) > 0 {
                            esq = esq[:len(esq) - 1]
                        }
                    case 'D':
                        if len(dir) > 0 {
                            dir = dir[:len(dir) - 1]
                        }
                    case '>':
                        if len(dir) > 0{
                            ultdir := dir[len(dir) - 1]
                            esq = append(esq, ultdir)
                            dir = dir[: len(dir) - 1]
                        }
                    case '<':
                        if len(esq) > 0{
                            ultesq := esq[len(esq) - 1]
                            dir = append(dir, ultesq)
                            esq = esq[: len(esq) - 1]
                        }
                    }
                }
            }

            resul := string(esq) + "|"
            for i := len(dir) - 1; i >= 0; i--{
                resul += string(dir[i])
            }
            fmt.Println(resul)
        
    }
    
}