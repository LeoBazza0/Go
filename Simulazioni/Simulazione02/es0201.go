package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(TrasformaParola(os.Args[i], i), " ")
	}
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	for i, r := range parola {
		switch {
		case posizione%2 != 0 && i%2 == 0:
			r = unicode.ToUpper(r)
		case posizione%2 == 0 && i%2 != 0:
			r = unicode.ToUpper(r)
		}
		parolaTrasformata += string(r)
	}
	return parolaTrasformata
}
