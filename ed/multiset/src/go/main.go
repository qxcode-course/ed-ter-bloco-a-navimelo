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
	newCapacity := ms.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		newData := make([]int, newCapacity)
		copy(newData, ms.data[:ms.size])
		ms.data = newData
		ms.capacity = newCapacity
}

func (ms *MultiSet) search (value int) int {
	l := 0
	r := ms.size - 1
	for l <= r {
		mid := (l + r) / 2
		if ms.data[mid] == value{
			return mid
		} else if ms.data[mid] < value {
			l = mid + 1 
		} else {
			r = mid - 1
		}
	}
	return -1
}

func (ms *MultiSet) insert (value int){
	if ms.size == ms.capacity {
		ms.expand()
	}
	idx :=  0
	for idx < ms.size && ms.data[idx] < value{
		idx++
	}

	for i := ms.size; i > idx; i--{
		ms.data[i] = ms.data[i - 1]
	}

	ms.data[idx] = value
	ms.size++
}

func (ms *MultiSet) Insert(value int) {
	ms.insert(value)
}

func (ms *MultiSet) erase (value int) bool{
	index := ms.search(value)
	if index != -1 {
		for i := index; i < ms.size - 1; i++{
			ms.data[i] = ms.data[i + 1]
		}
		ms.size--
		return true
	}
	return false
}

func (ms *MultiSet) Erase(value int) {
	ms.erase(value)
}	

func (ms *MultiSet) Contains (value int) bool {
	return ms.search(value) != -1
}

func (ms *MultiSet) Count(value int) int {
	index := ms.search (value)
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

func (ms *MultiSet) clear() {
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
			capacity, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(capacity)
		case "insert":
			for _, numstr := range args[1:] {
				num, _ := strconv.Atoi(numstr)
				ms.Insert(num)
			 }
		case "show":
			strValues := make ([]string, ms.size)
			for i := 0; i < ms.size; i++{
				strValues[i] = strconv.Itoa(ms.data[i])
			}
			fmt.Printf("[%s]\n", strings.Join(strValues, ", "))
		case "erase":
			num, _ := strconv.Atoi(args[1])
			if !ms.erase(num){
				fmt.Println("value not found")
			}
		case "contains":
			num, _ := strconv.Atoi(args[1])
			if ms.Contains(num){
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			num, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(num))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
