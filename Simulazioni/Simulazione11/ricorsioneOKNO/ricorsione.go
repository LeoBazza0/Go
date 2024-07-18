package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	// leggo prima il byte c e poi la stringa s
	c := []byte(os.Args[1])
	for i := 2; i < len(os.Args); i++ {
		s += os.Args[i]
	}
	fmt.Printf("%c è presente in %s in posizione: %d", c[0], s, isThere(c[0], s))
}

func isThere(c byte, s string) int {
	for i, k := range s {
		if k == rune(c) {
			return i + 1
		}
	}
	return -1
}

// è corretto ma non è ricorsivo e sinceramente no saprei com emetterlo altrimenti lol
