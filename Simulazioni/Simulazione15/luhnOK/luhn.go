package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//check presenza input
	if len(os.Args) == 1 {
		fmt.Println("no input")
		return
	}

	for _, number := range os.Args[1:] {
		fmt.Print(number)

		//controlla se dato valido
		if _, err := strconv.Atoi(number); len(number) != 16 || err != nil {
			fmt.Println(" dato non valido")
		} else {
			//controlla se luhn e stampa
			if isLuhn(number) {
				fmt.Println(" valido")
			} else {
				fmt.Println(" non valido")
			}
		}
	}
}

func isLuhn(number string) bool {
	sum := 0
	//algoritmo di Luhn
	//raddopia e somma cifre dispari
	for i := len(number) - 2; i >= 0; i = i - 2 {
		n := (int(number[i]) - '0') * 2
		if n > 9 {
			n -= 9
		}
		sum += n
	}
	//somma cifre pari
	for i := len(number) - 1; i >= 0; i = i - 2 {
		sum += int(number[i] - '0')
	}
	//controllo validità
	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}
