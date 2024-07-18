package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var tot int = 1
	stringa_inserita := os.Args
	//	stringaCheMiServe:=stringa_inserita[1:]
	sliceStringa := strings.Split(stringa_inserita[1], "")
	fmt.Println(sliceStringa)
	for i := 0; i < len(sliceStringa); i++ {
		num, err := strconv.Atoi(sliceStringa[i])
		if err == nil {
			tot *= num
		}
		if sliceStringa[i] == "#" {
			numno, _ := strconv.Atoi(sliceStringa[i-1])
			tot /= numno
		}
	}
	fmt.Println("il totale è: ", tot)
}

/*
1. il porgrama riceve u dati attraverso standard input
2. i dati di input sono inizialmente di tipo strings
3. i dati ricevuti da input devono essere tradotti da tipo stringa a tipo int ove possiibilie
4. i dati sono in serie
5. bisogna considerare sono il carattere che si legge e, nel caso sia "#" anche il carattere prima

*/
