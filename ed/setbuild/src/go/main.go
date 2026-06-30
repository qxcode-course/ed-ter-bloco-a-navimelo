package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct{
	data []int
	size int
	cap int
}

func NewSet(cap int) *Set{
	return &Set{
		data: make([]int, 0, cap),
		size: 0,
		cap: cap,
	}
}

func (s *Set) expand(){
	if s.cap == 0{
		s.cap = 1
		s.data = make([]int, s.cap)
	} else {
		s.cap *= 2
		newData := make([]int, s.size, s.cap)
		copy(newData, s.data)
		s.data = newData
	}
}

func (s *Set) Insert(valor int){
	for i := 0; i < s.size; i++{
		if s.data[i] == valor {
			return
		}
	}

	if s.size == s.cap {
		s.expand()
	}

	index := s.size
	for i := 0; i < s.size; i++{
		if valor < s.data[i]{
			index = i
			break
		}
	}

	s.data = append(s.data, 0)
	for i:= s.size; i> index; i--{
		s.data[i] = s.data[i - 1]
	}

	s.data[index] = valor
	s.size++
}

func (s *Set) String() string {
	return "[" + Join(s.data[0:s.size], ", ") + "]"
}

func (s *Set) Contains(valor int) bool{
	l, r := 0, s.size -1
	for l <= r{
		mid := (l + r)/2
		if s.data[mid] == valor{
			return true
		} else if s.data[mid] < valor{
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func (s *Set) Erase(valor int) error{
	l, r := 0, s.size -1
	index := -1

	for l <= r{
		mid := (l + r)/2
		if s.data[mid] == valor{
			index = mid
			break
		} else if s.data[mid] < valor{
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if index == -1{
		return fmt.Errorf("value not found")
	}

	for i := index; i < s.size - 1; i++{
		s.data[i] = s.data[i + 1]
	}
	s.size--
	s.data = s.data[:s.size]
	return nil
}

func (s *Set) Clear(){
	s.data = make([]int, 0, s.cap)
	s.size = 0
}

func Join(slice []int, sep string) string{
	if len(slice) == 0 {
		return ""
	}
	r := fmt.Sprintf("%d", slice[0])
	for _, valor := range slice[1:]{
		r += sep + fmt.Sprintf("%d", valor)
	}
	return r
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	v := NewSet(0)
	
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
			value, _ := strconv.Atoi(part)
			v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if err := v.Erase(value); err != nil{
				fmt.Println(err.Error())
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
