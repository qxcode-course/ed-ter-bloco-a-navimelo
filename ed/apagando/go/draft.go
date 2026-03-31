package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	
	fila := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var x int
	fmt.Scan(&x)
	
	saiufila := make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&saiufila[i])
			for j := 0; j <len(fila); j++ {
				if saiufila[i] == fila[j] {
					fila = append(fila[:j], fila[j+1:]...)
					break
				}
			}
	}
	
	saida := fmt.Sprintf("%v", fila)
    if saida == "[]" {
        fmt.Printf("N")
    } else {   
        fmt.Printf(saida[1 : len(saida)-1])
		fmt.Println(" ")
    }
}