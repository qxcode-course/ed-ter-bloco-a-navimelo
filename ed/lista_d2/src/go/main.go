//utilizando o código da lista_d1
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
	root *Node
}

func (n *Node) Next() *Node{
	if n.next == n.root{
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node{
	if n.prev == n.root{
		return nil
	}
	return n.prev
}

type LList struct{
	root *Node
	size int
}

func NewLList() *LList {
	sent := &Node{}
	sent.next = sent
	sent.prev = sent
	sent.root = sent

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
	novo := &Node{value: valor, prev: ll.root, next: ll.root.next, root: ll.root}
	ll.root.next.prev = novo
	ll.root.next = novo
	ll.size++
}

func (ll *LList) PushBack(valor int){
	novo := &Node{value: valor, prev: ll.root.prev, next: ll.root, root: ll.root}
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

func (ll *LList) Front() *Node{
	if ll.size == 0{
		return nil
	}
	return ll.root.next
}

func (ll *LList) Back() *Node{
	if ll.size == 0{
		return nil
	}
	return ll.root.prev
}

func (ll *LList) Search(value int) *Node{
	atual := ll.root.next
	for atual != ll.root{
		if atual.value == value{
			return atual
		}
		atual = atual.next
	}
	return nil
}

func (ll *LList) Insert(node *Node, value int){
	if node == nil{
		return
	}
	novo := &Node{value: value, prev: node.prev, next: node, root: ll.root}
	node.prev.next = novo
	node.prev = novo
	ll.size++
}

func (ll *LList) Remove(node *Node) *Node{
	if node == nil{
		return nil
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	ll.size--
	return node
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
		case "walk":
			 fmt.Print("[ ")
			 for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.value)
			 }
			 fmt.Print("]\n[ ")
			 for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.value)
			 }
			 fmt.Println("]")
		case "replace":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	node.value = newvalue
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "insert":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Insert(node, newvalue)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Remove(node)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
