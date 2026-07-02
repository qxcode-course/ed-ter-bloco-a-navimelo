package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	q := NewQueue[string]()
	equipes := "ABCDEFGHIJKLMNOP"

	for _, e := range equipes{
		q.Enqueue(string(e))
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 15 && scanner.Scan(); i++{
		linha := scanner.Text()

		var golX, golY int
		fmt.Sscan(linha, &golX, &golY)

		equipeX := q.Dequeue()
		equipeY := q.Dequeue()

		if golX > golY{
			q.Enqueue(equipeX)
		} else {
			q.Enqueue(equipeY)
		}
	}

	vencedor := q.Dequeue()
	fmt.Println(vencedor)
}
