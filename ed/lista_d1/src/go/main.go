package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct{
	value int
	next *Node
	prev *Node
}

type LList struct{
	root *Node
	size int
}

func NewLList() *LList {
	sent := &Node{}
	sent.next = sent
	sent.prev = sent

	return &LList{root: sent, size: 0}
}

func (ll *LList) Size() int{
	return ll.size
}

func (ll *LList) Clear(){
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) PushFront(valor int){
	novo := &Node{value: valor, prev: ll.root, next: ll.root.next}
	ll.root.next.prev = novo
	ll.root.next = novo
	ll.size++
}

func (ll *LList) PushBack(valor int){
	novo := &Node{value: valor, prev: ll.root.prev, next: ll.root}
	ll.root.prev.next = novo
	ll.root.prev = novo
	ll.size++
}

func (ll *LList) PopFront(){
	if ll.size == 0 {
		return
	}
	removido := ll.root.next
	ll.root.next = removido.next
	removido.next.prev = ll.root
	ll.size--
}

func (ll *LList) PopBack(){
	if ll.size == 0 {
		return
	}
	removido := ll.root.prev
	ll.root.prev = removido.prev
	removido.prev.next = ll.root
	ll.size--
}

func (ll *LList) String() string{
	var elems []string
	atual := ll.root.next

	for atual != ll.root{
		elems = append(elems, fmt.Sprint(atual.value))
		atual = atual.next
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

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
		case "show":
			 fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			 }
		case "push_front":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushFront(num)
			 }
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			 ll.PopFront()
		case "clear":
			 ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
