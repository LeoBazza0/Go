package main
import "fmt"
func main() {
    /*
    creare un ciclo che ti stampa un numero n di numeri e ogni volta, partendo da 0 aggiungi 2
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
