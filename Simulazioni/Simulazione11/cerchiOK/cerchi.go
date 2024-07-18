package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cerchio struct {
	nome         string
	x, y, raggio float64
}

func main() {
	var tangeOno, maggiOno string
	prec := Cerchio{"", 0, 0, 0}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ci := newCircle(line)
		if tocca(prec, ci) {
			tangeOno = "tangente,"
		} else {
			tangeOno = "non tangente,"
		}
		if maggiore(prec, ci) {
			maggiOno = "minore o uguale"
		} else {
			maggiOno = "maggiore"

		}
		fmt.Println(ci, tangeOno, maggiOno)
		prec = ci
	}

}

func newCircle(descr string) Cerchio {
	s := strings.Split(descr, " ")
	x, _ := strconv.ParseFloat(s[1], 64)
	y, _ := strconv.ParseFloat(s[2], 64)
	raggio, _ := strconv.ParseFloat(s[3], 64)
	return Cerchio{s[0], x, y, raggio}
}

func distanzaPunti(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func tocca(cerc1, cerc2 Cerchio) bool {
	if cerc1.raggio+cerc2.raggio == distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y) {
		return true

	} else if math.Abs(cerc1.raggio-cerc2.raggio) > distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)-math.Pow(10, -6) && math.Abs(cerc1.raggio-cerc2.raggio) < distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)+math.Pow(10, -6) {
		return true
	}
	return false
}

func maggiore(cerc1, cerc2 Cerchio) bool {
	return math.Pi*math.Pow(cerc1.raggio, 2) > math.Pi*math.Pow(cerc2.raggio, 2)
}
