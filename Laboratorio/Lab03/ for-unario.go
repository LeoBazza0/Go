package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    /*
    n è un numero random tra 0 e 10 + 1, lo stampa e poi aggiunge quel valore a t fin quando è minore del valore TARGET=20
    */
    const TARGET = 20
    const MAX = 10
    rand.Seed(time.Now().Unix())
    var n int

    t := 0
    for t < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        t += n
    }
    fmt.Println(t)
}
