package main
import (
    "fmt"
)

func podeSub(arr []rune, idx int, op rune, l int) bool{
    for i := 1; i <= l; i++{
        ant := idx - i
        if ant >= 0{
            if arr[ant] == op{
                return false
            }
        }
    }

    for i := 1; i <= l; i++{
        prox := idx + i
        if prox < len(arr){
            if arr[prox] == op{
                return false
            }
        }
    }
    return true
}

func resol(arr []rune, idx int, l int) bool{
    if idx == len(arr){
        return true
    }
    if arr[idx] != '.'{
        return resol(arr, idx + 1, l)
    }
    for i := 0; i <= l; i++{
        op := rune('0' + i)

        if podeSub(arr, idx, op, l){
            arr[idx] = op

            if resol(arr, idx + 1, l){
                return true
            }
        }
        arr[idx] = '.'
    }
    return false
}

func main(){
    var seq string
    var l int

    if _, err := fmt.Scan(&seq); err != nil{
        return
    }
    if _, err := fmt.Scan(&l); err != nil{
        return
    } 

    arr := []rune(seq)
    if resol(arr, 0, l){
        fmt.Println(string(arr))
    }

}