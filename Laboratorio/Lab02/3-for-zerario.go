package main
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
    /*il programma fa partire un ciclo che da numeri da 0 a K (20), li stampa con uno spazio in mezzo, e aumenta di un o il contatore c
    se stampa 0, il programma esce dal ciclo e stampa il valore del contatore all'ultimo ciclo */
    rand.Seed(time.Now().Unix())
    const K = 20
    var n int

    c := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        c++
    }
    fmt.Println(c)
}
