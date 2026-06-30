package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func aux (n int) int{
	if n < 0{
		return -n
	}
	return n
}

func getMen(vet []int) []int {
	var r []int
	for _, val := range vet{
		if val > 0{
			r = append(r, val)
		}
	}
	return r
}

func getCalmWomen(vet []int) []int {
	var r []int
	for _, val := range vet{
		if val < 0 && aux(val) < 10 {
			r = append(r, val)
		}
	}
	return r
}

func sortVet(vet []int) []int {
	r := make([]int, len(vet))
	copy(r, vet)
	
	sort.Ints(r)
	return r
}

func sortStress(vet []int) []int {
	r:= make([]int, len(vet))
	copy(r, vet)

	sort.Slice(r, func(i, j int) bool{
		return aux(r[i]) < aux(r[j])
	})
	return r
}

func reverse(vet []int) []int {
	var r [] int
	for i := len(vet) - 1; i >= 0; i--{
		r = append(r, vet[i])
	}
	return r
}

func unique(vet []int) []int {
	var r []int
	for _, val := range vet{
		existe := false
		for _, add := range r{
			if add == val {
				existe = true
				break
			}
		}
		if !existe{
			r = append(r, val)
		}
	}
	return r
}

func repeated(vet []int) []int {
	var r []int
	cont := make(map[int]int)

	for _, val := range vet{
		cont[val]++
		if cont[val] == 2{
			r = append(r, val)
		}
	}
	return r
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

