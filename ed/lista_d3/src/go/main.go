// usando o codigo lista_d2 como base
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}

func (ll *LList) String() string{
	var elems []string
	atual := ll.root.next

	for atual != ll.root{
		elems = append(elems, fmt.Sprint(atual.Value))
		atual = atual.next
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

func equals(lla, llb *LList) bool{
	atual1 := lla.root.next
	atual2 := llb.root.next

	for atual1 != lla.root && atual2 != llb.root{
		if atual1.Value != atual2.Value{
			return false
		}
		atual1 = atual1.next
		atual2 = atual2.next
	}
	return atual1 == lla.root && atual2 == llb.root
}

func addsorted(ll *LList, value int){
	atual := ll.root.next
	for atual != ll.root{
		if atual.Value >= value{
			ll.insertBefore(atual, value)
			return
		}
		atual = atual.next
	}
	ll.insertBefore(ll.root, value)
}

func reverse(ll *LList){
	atual := ll.root
	for {
		atual.next, atual.prev = atual.prev, atual.next
		atual = atual.prev
		if atual == ll.root{
			break
		}
	}
}

func merge(lla, llb *LList) *LList{
	merged := NewLList()
	atual1 := lla.root.next
	atual2 := llb.root.next

	for atual1 != lla.root && atual2 != llb.root{
		if atual1.Value <= atual2.Value{
			merged.PushBack(atual1.Value)
			atual1 = atual1.next
		} else {
			merged.PushBack(atual2.Value)
			atual2 = atual2.next
		}
	}

	for atual1 != lla.root{
		merged.PushBack(atual1.Value)
			atual1 = atual1.next
	}

	for atual2 != llb.root{
		merged.PushBack(atual2.Value)
			atual2 = atual2.next
	}

	return merged
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
			 	fmt.Println("iguais")
			} else {
			 	fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
			 	value, _ := strconv.Atoi(args[i])
			 	addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
