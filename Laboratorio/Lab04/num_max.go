package main

import (
	"fmt"
)

func main() {
	var n int
	g := 0
  count:=0

	fmt.Print("inserisci sequenza di 10 numeri: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		if n > g {
      if n!=g {
        count=0
      }
			g = n
    }
    if n==g {
count++
    }
	}
  fmt.Println("Il numero massimo è stato il ", g, "e si è ripetuto", count, "volte")
}
