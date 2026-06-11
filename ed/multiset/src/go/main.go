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
	if capacity < 0 {
		capacity = 0
	}
	return &MultiSet{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.size == ms.capacity {
		newCapacity := ms.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		newData := make([]int, ms.size, newCapacity)
		copy(newData, ms.data)
		ms.data = newData
		ms.capacity = newCapacity
	}
}

func (ms *MultiSet) search (value int) int {
	for i, v := range ms.data {
		if v == value {
			return i
		}
	}
	return -1
}

func (ms *MultiSet) insert (value int){
	if ms.size == ms.capacity {
		ms.expand()
	}
	ms.data = append(ms.data, value)
	ms.size++
}

func (ms *MultiSet) Insert(value int) {
	ms.insert(value)
}

func (ms *MultiSet) erase (value int) {
	index := ms.search(value)
	if index != -1 {
		ms.data = append(ms.data[:index], ms.data[index+1:]...)
		ms.size--
	}
}

func (ms *MultiSet) Erase(value int) {
	ms.erase(value)
}	

func (ms *MultiSet) contains (value int) bool {
	return ms.search(value) != -1
}

func (ms *MultiSet) Contains(value int) bool {
	return ms.contains(value)
}

func (ms *MultiSet) count (value int) int {
	count := 0
	for _, v := range ms.data {
		if v == value {
			count++
		}
	}
	return count
}

func (ms *MultiSet) Count(value int) int {
	return ms.count(value)
}

func (ms *MultiSet) unique() []int {
	uniqueValues := make(map[int]bool)
	for _, v := range ms.data {
		uniqueValues[v] = true
	}
	result := make([]int, 0, len(uniqueValues))
	for v := range uniqueValues {
		result = append(result, v)
	}
	return result
}

func (ms *MultiSet) Unique() []int {
	return ms.unique()
}

func (ms *MultiSet) clear() {
	ms.data = make([]int, 0, ms.capacity)
	ms.size = 0
}

func (ms *MultiSet) Clear() {
	ms.clear()
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
	_ = ms

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
			// for _, part := range args[1:] {
			// 	value, _ := strconv.Atoi(part)
			// }
		case "show":
		case "erase":
			// value, _ := strconv.Atoi(args[1])
		case "contains":
			// value, _ := strconv.Atoi(args[1])
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
