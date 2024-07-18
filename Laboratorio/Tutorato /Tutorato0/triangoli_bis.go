package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Inserisci tre numeri separati da spazi:")
	fmt.Scan(&a, &b, &c)
	if a > b+c || b > a+c || c > a+b {
		fmt.Println("la formula non funziona")
		return
	}
	if a == b && a == c {
		fmt.Println("Il triangolo è equilatero")
	} else if a == b || a == c || b == c {
		fmt.Println("Il triangolo è isoscele")
	} else {
		fmt.Println("Il triangolo è scaleno")
	}
}
