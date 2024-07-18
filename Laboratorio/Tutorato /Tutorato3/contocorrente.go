package main

import (
	"fmt"
)

func main() {
	var contocorr, addebito int
	fmt.Print("inserisci il conto corrente:")
	fmt.Scan(&contocorr)
	for contocorr > 0 {
		fmt.Print("inserisci addebito: ")
		fmt.Scan(&addebito)
		contocorr -= addebito
		fmt.Println("conto corrente attuale: ", contocorr)
	}
}
