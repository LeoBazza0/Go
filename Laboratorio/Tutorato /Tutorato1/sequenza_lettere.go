package main

import "fmt"

func main() {
	var previous, current int
	count := 0

	for previous <= current {
		fmt.Scan(&current)
		count++
		previous = current
	}
	fmt.Println("Numeri validi inseriti:", count)
}



leggo string --> slice string (solo lettere singole) --> converto in rune --> 
controllo che con quella prima vada bene --> faccio i calcoli con un for e la slice di rune 