package main

import "fmt"

func main() {
    /*
    scannerizzare un valore n al quale viene aggiunto un valore s che è creato dalla somma del valore s precedente e n.
    Si entra in un ciclo che fa inserire fino a 5 valori, dopodiché stampa la somma dei valori inseriti.
    */
    const K = 5
	var n int
	somma := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		somma = somma + n
	}
	fmt.Println(somma)
}
