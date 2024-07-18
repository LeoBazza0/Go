package main
import "fmt"
func main() {
	/*
dichiarare un valore intero,
se è maggioire di zero, si stampa positivo,
se è uguale a zero +, si stampa nullo,
se non è niente delle opzioni sopra, e quindi oer forza minore di zero, si stampa negativo
	*/

	var num int
	fmt.Print("un int: ")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else { // < 0
		fmt.Println("negativo")
	}
}
