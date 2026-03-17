package main
import "fmt"
func main() {
    var a string

    fmt.Scan(&a)

    var b int

    fmt.Scan(&b)

    if b < 12 {
        fmt.Println(a + " eh crianca")
    } else if b >= 12 && b < 18 {
        fmt.Println(a + " eh jovem")
    } else if b >= 18 && b < 65 {
        fmt.Println(a + " eh adulto")
    } else if b >= 65 && b < 1000{
        fmt.Println(a + " eh idoso")
    } else if b >= 1000{
        fmt.Println(a + " eh mumia")
    }

}
