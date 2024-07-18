package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Vocabolario struct {
	parola     string
	occorrenza int
}

func main() {
	var testo, line string
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("indicare nome file e iniziale")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		testo += line
	}
	carattere := os.Args[2]
	fmt.Println(carattere)
	String(creaVoc(testo))
}

func creaVoc(testo string) (sliceVocabolario []Vocabolario) {
	sliceTesto := strings.Fields(testo)
	sort.Strings(sliceTesto)
	for i := 0; i < len(sliceTesto); i++ {
		sliceVocabolario = append(sliceVocabolario, Vocabolario{sliceTesto[i], strings.Count(testo, sliceTesto[i])})
		if strings.Count(testo, sliceTesto[i]) > 1 {
			testo1, testo2, _ := strings.Cut(testo, sliceTesto[i])
			testo = testo1 + testo2
		}
	}
	return
}

/*
func ([]Vocabolario) String() { //dovrei fare un metodo ma non sono capace

		for i := 0; i < len(sliceVocabolario); i++ {
			fmt.Println(sliceVocabolario[i].parola, ":", sliceVocabolario[i].occorrenza)
		}
	}
*/
func String(sliceVocabolario []Vocabolario) { //dovrei fare un metodo ma non sono capace
	for i := 0; i < len(sliceVocabolario); i++ {
		fmt.Println(sliceVocabolario[i].parola, ":", sliceVocabolario[i].occorrenza)
	}
}
