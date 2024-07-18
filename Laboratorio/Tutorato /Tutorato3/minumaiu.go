package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	/*
			stringaIniz := os.Args[1:]
		//	fmt.Println(stringaIniz)
			joined := strings.Join(stringaIniz, "")
			fmt.Println("joined: ", joined)
			//manca tutto il pezzo dopo
	*/
	for _, parola := range os.Args[1:] {
		for i := 0; i < len(parola)-1; i++ {
			if unicode.IsLower(rune(parola[i])) {
				if !unicode.IsUpper(rune(parola[i+1])) {
					fmt.Println("sequenza non valida")
					os.Exit(1)
				}
			} else if unicode.IsUpper(rune(parola[i])) {
				if !unicode.IsLower(rune(parola[i+1])) {
					fmt.Println("sequenza non valida")
					os.Exit(2)
				}
			}
		}
		fmt.Println("sequenza valida")
	}
}
