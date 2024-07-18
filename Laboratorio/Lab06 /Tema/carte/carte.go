package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Carta struct {
	valore, seme string
}

const semi = 4
const valori = 13
const cartetot = 52

func main() {
	var n int
	fmt.Print("Inserisci un numero di una carta [0-52): ")
	fmt.Scan(&n)
	cartastampata, bool := carta(n)
	if bool == false {
		fmt.Println("argomento non compreso nell'intervallo", cartastampata)
	} else {
		fmt.Println(cartastampata)
	}
	cartaRandom, _ :=carta(estraiCarta())
	fmt.Println ("carta random: ", cartaRandom)
	fmt.Println ("quattro carte: ", dai4Carte())
}

func carta(n int) (carta Carta, bool bool) {
	type Carta struct {
		valore, seme string
	}
	if n < 0 || n >= 52 {
		carta.valore = "//"
		carta.seme = "//"
		bool = false
		return carta, bool
	} else {
		if n >= 0 && n <= 12 {
			carta.seme = "cuori"
		} else if n >= 13 && n <= 25 {
			carta.seme = "quadri"
		} else if n >= 26 && n <= 38 {
			carta.seme = "fiori"
		} else if n >= 39 && n <= 51 {
			carta.seme = "picche"
		}
		for i := 0; i < 4; i++ {
			if n == 0+13*i {
				carta.valore = "A"
			} else if n == 1+13*i {
				carta.valore = "2"
			} else if n == 2+13*i {
				carta.valore = "3"
			} else if n == 3+13*i {
				carta.valore = "4"
			} else if n == 4+13*i {
				carta.valore = "5"
			} else if n == 5+13*i {
				carta.valore = "6"
			} else if n == 6+13*i {
				carta.valore = "7"
			} else if n == 7+13*i {
				carta.valore = "8"
			} else if n == 8+13*i {
				carta.valore = "9"
			} else if n == 9+13*i {
				carta.valore = "10"
			} else if n == 10+13*i {
				carta.valore = "J"
			} else if n == 11+13*i {
				carta.valore = "Q"
			} else if n == 12+13*i {
				carta.valore = "K"
			}
		}
		return carta, true
	}
}

func estraiCarta () int {
	 rand.Seed(time.Now().Unix())
	return rand.Intn(53)
}

func dai4Carte () []Carta {
	var	carte4 []Carta
	rand.Seed(time.Now().Unix())
	for i:=0 ; i<4; i++ {
		cartan, _:=carta(rand.Intn(53))
		carte4=append(carte4, cartan)
	}
	return carte4
}
