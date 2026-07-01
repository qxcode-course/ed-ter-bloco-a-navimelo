package main
import (
    "bufio"
    "fmt"
    "os"
)

type cursor struct{
    esq string
    dir string
}

func editor(comandos string) string {
    var esq []rune
    var dir []rune

    var histZ []cursor
    var histY []cursor

    for _, c := range comandos {
        if c == 'Z'{
            if len(histZ) > 0{
                atual := cursor{esq: string(esq), dir: string(dir)}
                histY = append(histY, atual)

                ult := histZ[len(histZ) - 1]
                histZ = histZ[: len(histZ) - 1]

                esq = []rune (ult.esq)
                dir = []rune (ult.dir)
            } else if c == 'Y' {
                if len(histY) > 0{
                atual := cursor{esq: string(esq), dir: string(dir)}
                histZ = append(histZ, atual)

                ult := histY[len(histY) - 1]
                histY = histY[: len(histY) - 1]

                esq = []rune (ult.esq)
                dir = []rune (ult.dir)
            }
        } else {
            atual := cursor{esq: string(esq), dir: string(dir)}
            histZ = append(histZ, atual)
            histY = nil

            if(c >= 'a' && c <= 'z') || c == '-'{
                esq = append(esq, c)
            } else {
                switch c {
                case 'R':
                    esq = append(esq, '\n')
                case 'B':
                    if len(esq) > 0{
                        esq = esq[:len(esq) -1]
                    }
                case 'D':
                    if len(dir) > 0{
                        dir = dir[:len(dir) -1]
                    }
                case '>':
                    if len(dir) > 0{

                        
                        dir = dir[:len(dir) -1]
                    }
                case '<':
                }
            }
        }
    }
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)

}