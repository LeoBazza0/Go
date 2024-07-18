package main

import (
	"fmt"
	"os"
)

func main() {
	var parola, prec string

	for {
		fmt.Println("inserici una parola")
		fmt.Scan(&parola)
		if verifica(parola) == false || len(parola) != len(prec) {
			fmt.Println("errore")
		} else if len(os.Args) == 1 {
			break
		}

		prec = parola
	}

}
func verifica(s string) bool {
	for _, k := range s {
		if k < 'a' || k > 'z' {
			return false
		}
	}
	return true
}

func verificaLettere(parola, prec string) (letteraOld, letteraNew, parolaNew string) {

}
