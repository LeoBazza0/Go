package main

import (
	"fmt"
	"os"
	"strconv"
)

type Cisterna struct {
	larghezza, lunghezza, altezza float64
	dovearriva                    int
}

func main() {
	var lt int
	var cist Cisterna
	var err error
	la, err1 := strconv.ParseFloat(os.Args[1], 64)
	lu, err2 := strconv.ParseFloat(os.Args[2], 64)
	al, err3 := strconv.ParseFloat(os.Args[3], 64)
	if la < 0 || lu < 0 || al < 0 {
		fmt.Println("tutti gli argomenti devono essere >0")
		os.Exit(0)
	} else if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("tutti gli argomenti devono essere numerici")
		os.Exit(0)
	} else {
		cist = Cisterna{la, lu, al, 0}
	}
	for {
		fmt.Scan(&lt)
		if lt == 9999 {
			os.Exit(0)
		} else {
			err = variazione(&cist, lt)
		}
		fmt.Println(cist, err)
	}
}

// TODO
func variazione(cisterna *Cisterna, lt int) (err error) {
	if lt > 0 && volume(*cisterna, cisterna.dovearriva)+float64(lt)/(cisterna.lunghezza*cisterna.larghezza) < volume(*cisterna, int(cisterna.altezza)) {
		//riempi la cisterna con lt
		return nil
	}
	return
}

func volume(cisterna Cisterna, livello int) (vol float64) {
	return cisterna.larghezza * cisterna.lunghezza * float64(livello)
}
