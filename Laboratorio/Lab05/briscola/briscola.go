package main

import "fmt"

func puntiCarta(s string) int {
	var punti int
	for i := 0; i < len(s); i++ {
		if i > 3 {
			return 100
		}
		if s[i] != 'K' && s[i] != 'Q' && s[i] != 'J' && s[i] != 'A' && s[i] != '2' && s[i] != '3' && s[i] != '4' && s[i] != '5' && s[i] != '6' && s[i] != '7' {
			punti = -1
			break
		}
		switch s[i] {
		case 'A':
			punti += 11
		case '3':
			punti += 10
		case 'K':
			punti += 4
		case 'Q':
			punti += 3
		case 'J':
			punti += 2
		default:
			punti += 0
		}
	}
	return punti
}

func main() {
	var s string
	fmt.Print("Inserire le tre carte: ")
	fmt.Scan(&s)
	if puntiCarta(s) == 100 {
		fmt.Print("troppe carte")
	} else {
		fmt.Print("la tua mano ", s, ": ", puntiCarta(s), "punti")
	}
}
