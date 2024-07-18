package main
import "fmt"
func main() {
    /*
	 creare un ciclo che stampa u numero n astierischi
   fare uno fmt.scan per leggere un numero n e poi creare un ciclo per stampare gli asterischi
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
