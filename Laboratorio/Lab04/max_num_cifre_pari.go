package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nString string
	var count, max int
	fmt.Println("inserisci una serie di numeri")
	for nString != "000" {
		fmt.Scan(&nString)
		_, err := strconv.Atoi(nString)
		if err != nil {
			fmt.Println("Il numero che hai inserito non è un numero intero")
		}
		count = 0
		for _, cifra := range nString {
			if (cifra-'0')%2 == 0 {
				count++
			}
		}
		if count > max {
			max = count
		}
}
fmt.Println("Il numero con più cifre pari massimo ha", max, "cifre pari")
}
//non funziona per ora :/
