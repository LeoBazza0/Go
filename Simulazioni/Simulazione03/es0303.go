package main

import (
	"fmt"
	"os"
)

func main() {
	r := os.Args[1]
	s := os.Args[2]
	count := occorrenze(r, s)
	fmt.Printf("occorrenze di %s in %s \n %d", r, s, count)
}

func occorrenze(r, s string) (count int) {
	for _, k := range s {
		if string(k) == r {
			count++
		}
	}
	return count
}
