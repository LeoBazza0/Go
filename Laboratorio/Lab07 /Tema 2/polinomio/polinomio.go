package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Scrivere un programma polinomio.go che legge da standard input una riga che contiene dei numeri a, b, ....
Il programma calcola il valore in x del polinomio
a + bx + cx^2 + dx^3 ....
corrispondente alla sequenza dei numeri letti tranne l'ultimo, dove il valore di x è l'ultimo valore della sequenza.
Ad esempio,
3 2 0 7 5
corrisponde al polinomio 3 + 2x + 7x^3 da valutare per x = 5
*/
/*
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	dati := strings.Split(input, " ")
	fmt.Println(valutaPolinomio(dati[:len(dati)-1], dati[len(dati)-1]))
}

func valutaPolinomio(polinomio []string, x string) (t float64) {
	xInt, _ := strconv.ParseFloat(x, 64)
	for i, s := range polinomio {
		value, _ := strconv.ParseFloat(s, 64)
		t += value * math.Pow(xInt, float64(i))
	}
	return
}
*/

func main() {
	var sliceF []float64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	sliceS := strings.Split(line, " ")
	for i, _ := range sliceS {
		f, _ := strconv.ParseFloat(sliceS[i], 64)
		sliceF = append(sliceF, f)
	}
	fmt.Println(result(sliceF))
}

func result(sliceF []float64) (ris float64) {
	for i := 0; i < len(sliceF)-1; i++ {
		ris += sliceF[i] * math.Pow(sliceF[len(sliceF)-1], float64(i))
	}
	return ris
}

//func Pow(x, y float64) float64
