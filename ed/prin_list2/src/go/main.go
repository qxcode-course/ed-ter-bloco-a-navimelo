package main

import (
	"container/list"
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *list.List, sword *list.Element) string {
	var sb strings.Builder
	sb.WriteString("[ ")

	for n := l.Front(); n != nil; n = n.Next(){
		valor := n.Value.(int)

		if n == sword{
			if valor > 0{
				sb.WriteString(fmt.Sprintf("%d> ", valor))
			} else {
				sb.WriteString(fmt.Sprintf("<%d ", valor))
			}
		} else {
			sb.WriteString(fmt.Sprintf("%d ", valor))
		}
	}

	sb.WriteString("]")
	return sb.String()
}

// move para frente na lista circular
func Next(l *list.List, it *list.Element) *list.Element {
	n := it.Next()
	if n == nil{
		return l.Front()
	}
	return n
}

// move para tras na lista circular
func Prev(l *list.List, it *list.Element) *list.Element {
	p := it.Prev()
	if p == nil{
		return l.Back()
	}
	return p
}

func main() {
	var qtd, chosen, fase int
	fmt.Scan(&qtd, &chosen, &fase)
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i * fase)
		fase = -fase
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		if sword.Value.(int) > 0 {
			l.Remove(Next(l, sword))
			sword = Next(l, sword)
		} else {
			l.Remove(Prev(l, sword))
			sword = Prev(l, sword)
		}
	}
	fmt.Println(ToStr(l, sword))
}
