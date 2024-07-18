package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s [10]int
	var n int
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(os.Args[1+i])
		s[i] = num
	}
	fmt.Println("Inserisci un numero:")
	fmt.Scan(&n)
	p := -1
	for i, z := range s {
		if z == n {
			p = i
		}
	}
	fmt.Println(p)
}
