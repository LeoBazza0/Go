package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var m map[string]capoluoghi
var text, sigla, sigla2 string
var textslice []string

type capoluoghi struct {
	nome, sigla, regione string
	superficie           int
}

// nome, sigla, regione, (popolazione),superficie, (densita), (altitudine)
func main() {
	var cap capoluoghi
	m = make(map[string]capoluoghi)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToUpper(line)
		lineslice := strings.Split(line, ",")
		for i := 0; i < 8; i++ {
			sigla = lineslice[1]
			cap.sigla = sigla
			cap.nome = lineslice[0]
			cap.regione = lineslice[2]
			superficie, _ := strconv.Atoi(lineslice[4])
			cap.superficie = superficie
			m[sigla] = cap
		}
		text += line
	}
	textslice = strings.Split(text, ",")
	//fmt.Println(m)
	for j := 1; j < len(os.Args); j++ {
		sigla := os.Args[j]
		fmt.Println(sigla, ":", m[sigla])
	}
	fmt.Println(m)
}
