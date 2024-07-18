package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("inserisci un numero ≥ a zero: ")
	fmt.Scan(&numero)
	numToText(numero)
}

func numToText(num int) {
	var cifra string
	var res string
	numeri := make(map[int]string)
	for num > 0 {
		digit := num % 10
		if _, ok := numeri[digit]; !ok {
			fmt.Print ("parola per ", digit, "?")
			fmt.Scan(&cifra)
			numeri[digit] = cifra
		}
		num = num / 10
		res = numeri[digit] + "-" + res
	}
	fmt.Println(res)
}
