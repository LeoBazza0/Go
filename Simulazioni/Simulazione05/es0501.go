package main

import (
	"fmt"
)

type Persona struct {
	nome    string
	altezza float64
	eta     int
}

func main() {
	persone := []Persona{Persona{"Alice", 150, 12}, Persona{"Pincopanco", 170, 15}, Persona{"Pancopinco", 170, 15}}
	fmt.Print(persone)
	mangiaFungo(&persone[0])
	fmt.Println(entraTana(persone))
}

func mangiaFungo(p *Persona) {
	p.altezza /= 100
}

func entraTana(persone []Persona) (sliceET []Persona) {
	for _, k := range persone {
		if k.altezza < 3 {
			sliceET = append(sliceET, k)
		}
	}
	return sliceET
}
