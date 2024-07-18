package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func moda(serie []float64) float64 {
	var mapNum map[float64]int
	var moda float64
	var numero int
	mapNum = make(map[float64]int)
	for i := 0; i < len(serie); i++ {

	}
	for i := 0; i < len(serie); i++ {
		mapNum[serie[i]]++
		if mapNum[serie[i]] > numero {
			numero = mapNum[serie[i]]
			moda = serie[i]
		}
	}
	//	fmt.Println(mapNum)
	return moda
}

func sortfr(serie1 []float64) []float64 {
	sort.Float64s(serie1)
	if !sort.Float64sAreSorted(serie1) {
		sortfr(serie1)
	}
	return serie1
}

func mediana(serie []float64) float64 {
	if len(serie)%2 == 0 {
		return (serie[(len(serie)/2)-1] + serie[(len(serie)/2)]) / float64(2)
	} else {
		return serie[(len(serie) / 2)]
	}
}

func main() {
	var serie []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numero, _ := strconv.ParseFloat(scanner.Text(), 64)
		serie = append(serie, numero)
	}
	serie = sortfr(serie)
	//	fmt.Printf("moda: %f\nmediana: %f", moda(serie), mediana(serie))
	fmt.Println("moda:", moda(serie))
	fmt.Println("mediana:", mediana(serie))

}

/*
func moda(serie []float64) float64 {
	serie = sortfr(serie)

	var mapNum map[float64]int
	var numiniz int
	var moda float64
	mapNum = make(map[float64]int)
	for i := 0; i < len(serie); i++ {
		mapNum[serie[i]]++
	}
	for i, k := range mapNum {
		//	fmt.Println("k ", k, "; i", i, "numiniz", numiniz, "moda", moda)
		if k > numiniz {
			numiniz = k
			moda = i
		}
	}
	return moda
}
*/
