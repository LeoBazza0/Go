package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta         string
	ascissa, ordinata float64
}

func main() {
	tragitto := NuovoTragitto()
	lunghezza := Lunghezza(tragitto)
	fmt.Printf("Lunghezza percorso: %.3f\n", lunghezza)
	fmt.Printf("Punto oltre metà: %s\n", String(PuntoOltreMeta(tragitto, lunghezza/2.0)))
}

func NuovoTragitto() (tragitto []Punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dati := strings.Split(line, ";")
		flo1, _ := strconv.ParseFloat(dati[1], 64)
		flo2, _ := strconv.ParseFloat(dati[2], 64)
		tragitto = append(tragitto, Punto{dati[0], flo1, flo2})
	}
	return
}

func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p1.ascissa-p2.ascissa, 2) + math.Pow(p1.ordinata-p2.ordinata, 2))
}

func String(p Punto) string {
	return fmt.Sprintf("%s = (%.1f, %.1f)", p.etichetta, p.ascissa, p.ordinata)
}

func Lunghezza(tragitto []Punto) float64 {
	var lunghezzaTot float64
	for i := 0; i < len(tragitto)-1; i++ {
		lunghezzaTot += Distanza(tragitto[i], tragitto[i+1])
	}
	return lunghezzaTot
}

func PuntoOltreMeta(tragitto []Punto, meta float64) Punto {
	var lunghezzaTot float64
	var i int
	for i = 0; i < len(tragitto)-1; i++ {
		lunghezzaTot += Distanza(tragitto[i], tragitto[i+1])
		if lunghezzaTot > meta {
			break
		}
	}
	return tragitto[i+1]
}
