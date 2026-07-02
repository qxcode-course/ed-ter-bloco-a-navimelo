package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var sb strings.Builder
	sb.WriteString("[ ")

	for n := l.root.next; n != l.root; n = n.next {
		sb.WriteString(fmt.Sprint(n.Value))
		if n == sword{
			sb.WriteString(">")
		}
		sb.WriteString(" ")
	}
	sb.WriteString("]")
	return sb.String()
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	next := it.next
	if next == l.root{
		next = next.next
	}
	return next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
