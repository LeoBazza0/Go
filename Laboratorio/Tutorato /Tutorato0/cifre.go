package main

import "fmt"

func main() {
	var n, i int
	var stop bool
	var a [10]int
	fmt.Print("Scrivi un numero intero >=0: ")
	fmt.Scan(&n)
	for n > 0 && !stop {
		c := n % 10
		n /= 10
		fmt.Println(c)
		a[i] = c
		i++
		if i >= 10 {
			stop = true
		}
	}
	if stop {
		fmt.Println("numero troppo grande")
	} else {
		fmt.Println(a)
	}
}
