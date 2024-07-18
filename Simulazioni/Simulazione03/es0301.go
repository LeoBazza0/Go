package main

import (
	"fmt"
	"os"
	"strconv"
)

type Orario struct {
	s, m, h int
}

func main() {
	s, errs := strconv.Atoi(os.Args[1])
	m, errm := strconv.Atoi(os.Args[2])
	h, errh := strconv.Atoi(os.Args[3])
	ss, _ := strconv.Atoi(os.Args[4])
	bool, ora := newOrario(s, m, h)
	if errs != nil || errm != nil || errh != nil || bool == false {
		fmt.Print("parametri non validi")
		os.Exit(0)
	} else {
		fmt.Printf("%d:%d:%d\n", ora.s, ora.m, ora.h)
	}
	for i := 0; i < ss; i++ {
		tic(&ora)
	}
	fmt.Printf("%d:%d:%d", ora.s, ora.m, ora.h)
}

func newOrario(s, m, h int) (bool, Orario) {
	if s <= 60 && s >= 0 && m <= 60 && m >= 0 && h <= 23 && h >= 0 || (h == 24 && m == 0 && s == 0) {
		ora := Orario{s, m, h}
		return true, ora
	} else {
		ora := Orario{0, 0, 0}
		return false, ora
	}

}

func tic(orario *Orario) {
	orario.s++
	if orario.s >= 60 {
		orario.m++
		orario.s = 0
	}
	if orario.m >= 60 {
		orario.h++
		orario.m = 0
	}
	if orario.h >= 24 {
		orario.h -= 24
	}
	return
}
