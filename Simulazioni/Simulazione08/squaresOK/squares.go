package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("too few arguments")
		os.Exit(0)
	}
	numero, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("required integer value")
		os.Exit(0)
	} else {
		fmt.Println(NestedSquares(numero))
		os.Exit(0)
	}
}

func NestedSquares(n int) string {
	if n < 5 {
		return ""
	}
	primaRiga := strings.Repeat("*", n) + "\n"
	UltimaRiga := strings.Repeat("*", n)
	secondaPenultimariga := "*" + strings.Repeat(" ", n-2) + "*" + "\n"
	altreRighe1 := "*" + " " + strings.Repeat("*", n-4) + " " + "*" + "\n"
	if n == 5 {
		return primaRiga + secondaPenultimariga + altreRighe1 + secondaPenultimariga + UltimaRiga
	} else {
		altreRighe2 := "*" + " " + "*" + strings.Repeat(" ", n-6) + "*" + " " + "*" + "\n"
		return primaRiga + secondaPenultimariga + altreRighe1 + strings.Repeat(altreRighe2, n-6) + altreRighe1 + secondaPenultimariga + UltimaRiga
	}

}

/*
Scrivere un programma 'squares.go' dotato di:

- una funzione
	func NestedSquares(n int) string
che costruisce e restituisce una stringa corrispondente ad una figura
formata da un quadrato di '*' di dimensione n (con n > 4), che racchiude
un quadrato di '*' di dimensione 'n - 4', separato da quello esterno da una
"cornice" di spazi di spessore unitario (si vedano gli esempi).
La stringa restituita NON deve terminare con '\n'.
Se n < 5, la funzione restituisce la stringa vuota

- una funzione main() che legge da linea di comando un intero
e produce su standard output la figura di quadrati della misura fornita.
Se non ci sono dati sulla linea di comando, il programma stampa il messaggio
"too few arguments" e termina.
Se il valore passato da linea di comando non è un intero, il programma stampa il messaggio
"required integer value" e termina.

Esempi
------
per n = 5
*****
*   *
* * *
*   *
*****
per n = 6
******
*    *
* ** *
* ** *
*    *
******
per n = 7
*******
*     *
* *** *
* * * *
* *** *
*     *
*******
*/
