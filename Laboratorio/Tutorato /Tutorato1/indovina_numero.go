package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const lower int = 1
	const upper int = 100
	var numero_inserito int
	numero_da_indovinare := numeroDaIndovinare(lower, upper)
	//fmt.Println(numero_da_indovinare)
	for {
		fmt.Print("inserisci un numero: ")
		fmt.Scan(&numero_inserito)
		if numero_inserito < numero_da_indovinare {
			fmt.Println("+")
		} else if numero_inserito > numero_da_indovinare {
			fmt.Println("-")
		} else {
			fmt.Println("Hai indovinato!")
			break
		}
	}
}

func numeroDaIndovinare(lower, upper int) (numero_da_indovinare int) {
	numero_da_indovinare = rand.Intn(upper + 1)
	return
}

//1. il programma riceve i numeri da confrontare con il numero random da stadnard input,
// il numero random è calcolato con la funzione numeroDaIndovinare
//2. int
//3. i numeri sono sparsi
//4. i dati inseriti dall'utente sono pronti per essere utilizzati
//5. è necessario avere memorizzato solo l'ultimo numero, in modo da confrontarlo con il numero generato random
//6. quando l'utente indovina il numero
