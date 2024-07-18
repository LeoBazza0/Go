package main

import "fmt"

func main() {
	var s [10]int
	var n int
	fmt.Println("Inserisci 10 numeri:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Println("Inserisci un numero:")
	fmt.Scan(&n)
	p := -1

	for i := 9; i >= 0; i-- {
		if s[i] == n {
			p = i
		}
	}

	fmt.Println(p)
}
