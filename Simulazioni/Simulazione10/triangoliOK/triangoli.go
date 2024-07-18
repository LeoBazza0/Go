package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Punto struct {
	x, y float64
}
type Triangolo struct {
	P1, P2, P3 Punto
}

func main() {
	var punti []Punto
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		x, _ := strconv.ParseFloat(s[0], 64)
		y, _ := strconv.ParseFloat(s[1], 64)
		p := Punto{x, y}
		punti = append(punti, p)
	}
	fmt.Println(lunghezzeLati(punti[0], punti[1], punti[2]))
	triang, err := newTriangolo(punti[0], punti[1], punti[2])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("triangolo", triang)
		fmt.Println("tipo", tipoTriangolo(triang))
	}
}

func lunghezzeLati(A, B, C Punto) (lunghezze [3]float64) {
	//che restituisce una array di tre elementi corrispondenti, nell'ordine, alle distanze  A-B, B-C e C-A
	lunghezze[0] = distanza(A, B)
	lunghezze[1] = distanza(B, C)
	lunghezze[2] = distanza(C, A)
	return
}

func distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func newTriangolo(A, B, C Punto) (Triangolo, error) {
	if distanza(A, B)+distanza(B, C) > distanza(C, A) && distanza(A, B)+distanza(A, C) > distanza(C, B) && distanza(A, C)+distanza(A, B) > distanza(C, B) {
		return Triangolo{A, B, C}, nil
	}
	err := errors.New("non è un triangolo")
	return Triangolo{A, B, C}, err
}

func tipoTriangolo(triangolo Triangolo) int {
	delta := 1e-6
	dist := lunghezzeLati(triangolo.P1, triangolo.P2, triangolo.P3)
	sort.Float64s(dist[:])
	if math.Abs(dist[0]-dist[1]) < delta && math.Abs(dist[1]-dist[2]) < delta {
		return 3
	} else if math.Abs(dist[0]-dist[1]) > delta && math.Abs(dist[1]-dist[2]) > delta {
		return 0
	} else {
		return 2
	}
}

/*
	- 0 se il triangolo è scaleno

	- 2 se il triangolo è isoscele

	- 3 se il triangolo è equilatero
*/
