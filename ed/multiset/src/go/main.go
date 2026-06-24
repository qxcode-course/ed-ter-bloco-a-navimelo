package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data: make([]int, 0 , capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity == 0 {
		ms.capacity = 1
		ms.data = make([]int, ms.capacity)
	} else {
		ms.capacity *= 2
		newData := make([]int, ms.size, ms.capacity)
		copy(newData, ms.data)
		ms.data = newData
	}
}

func (ms *MultiSet) Insert (value int){
	if ms.size == ms.capacity{
		ms.expand()
	}
	index := ms.size

	for i:= 0; i < ms.size; i++ {
		if value < ms.data[i] {
			index = i
			break
		}
	}

	ms.data = append(ms.data, 0)
	for i:= ms.size; i > index; i-- {
		ms.data[i] = ms.data[i - 1]
	}

	ms.data[index] = value
	ms.size++
}

func (ms *MultiSet) String() string{
	return "[" + Join(ms.data[0: ms.size], ", ") + "]"
}

func (ms *MultiSet) Contains (value int) bool {
	l, r := 0, ms.size - 1
	for l <= r {
		mid := (l + r)/2
		if ms.data[mid] == value {
			return true
		} else if ms.data[mid] < value {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func (ms *MultiSet) Erase(value int) error{
	l, r := 0, ms.size - 1
	index := -1
	for l <= r {
		mid := (l + r)/2
		if ms.data[mid] == value {
			index = mid
			break
		} else if ms.data[mid] < value {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if index == -1 {
		return fmt.Errorf("value not found")
	}
	for i := index; i < ms.size - 1; i++{
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	ms.data = ms.data[:ms.size]

	return nil
}	

func (ms *MultiSet) Count(value int) int {
	l, r := 0, ms.size - 1
	index := -1
	for l <= r {
		mid := (l + r)/2
		if ms.data[mid] == value{
			index = mid
			break
		} else if ms.data[mid] < value{
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if index == -1 {
		return 0
	}
	count := 1
	for i := index - 1; i >= 0 && ms.data[i] == value; i--{
		count ++
	}
	for i := index + 1; i < ms.size && ms.data[i] == value; i++{
		count ++
	}
	return count 
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0{
		return 0
	}
	dif := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i - 1]{
			dif++
		}
	}
	return dif
}

func (ms *MultiSet) Clear() {
	ms.data = make([]int, 0, ms.capacity)
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			 }
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if err := ms.Erase(value); err != nil{
				fmt.Println(err.Error())
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
