package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	VOCALI string = "aeiouAEIOU"
	CIFRE  string = "0123456789"
)
const (
	S1 string = "c"
	S2 string = "sc"
)

var parola string

func main() {
	fmt.Print("Inserisci una parola: ")
	fmt.Scan(&parola)

	/*
		fmt.Print("Inserisci S1 (un carattere): ")
		fmt.Scan(&S1)
		for {
			if len(S1) != 1 {
				fmt.Print("Stringa troppo lunga \n Inserisci S1 (un carattere): ")
				fmt.Scan(&S1)
			} else {
				break
			}
		} //controllo che effettivamente sia lungo un carattere
		fmt.Print("Inserisci S2 (più di un carattere): ")
		fmt.Scan(&S2)
		for {
			if len(S2) == 1 {
				fmt.Print("Stringa troppo corta \nInserisci S2 (più di un carattere): ")
				fmt.Scan(&S2)
			} else {
				break
			}
		} //controllo che effettivamente sia più lungo di un carattere
	*/

	usaStandard(parola)
}

func usaStandard(string) {
	fmt.Print("1. ")
	if strings.Contains(parola, S2) {
		fmt.Println(parola, "contiene", S2)
	} else {
		fmt.Println(parola, "non contiene", S2)
	}

	fmt.Print("2. ")
	if strings.ContainsAny(parola, VOCALI) {
		fmt.Println(parola, "contiene vocali")
	} else {
		fmt.Println(parola, "non contiene vocali")
	}

	fmt.Print("3. ")
	fmt.Println(parola, "ha", strings.Count(parola, S1), S1)

	fmt.Print("4. ")
	IndexPrimaVocale := strings.IndexAny(parola, VOCALI)
	IndexUltimaVocale := strings.LastIndexAny(parola, VOCALI)
	fmt.Println("la prima vocale di", parola, "è in posizione", IndexPrimaVocale, "l'ultima in posizione", IndexUltimaVocale)

	fmt.Print("5. ")
	fmt.Println(strings.Repeat((S2), 3))

	fmt.Println("6. ")
	//non c'è l'istruzione

	fmt.Print("7. ")
	var s string
	for i := 0; i < len(parola); i++ {
		if strings.Contains(CIFRE, string(parola[i])) {
			s += string(parola[i])
		}
	}
	sint, err := strconv.Atoi(s)
	if err != nil {
		fmt.Print("c'è stato un problema")
	}
	fmt.Println("int", sint)

  fmt.Print("8. ")
  f:=float64(sint)*0.001
  fmt.Printf("float64 %.6f\n", f)
}
