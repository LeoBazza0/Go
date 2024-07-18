package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var line string
	numero, _ := strconv.Atoi(os.Args[1])
	file, _ := os.Open(os.Args[2])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line += scanner.Text()
	}
	lineSlice := strings.Fields(line)
	fmt.Println(dividi(numero, lineSlice))
}

func dividi(numero int, lineSlice []string) (righe []string) {
	var rigaper string
	for i := 0; i < len(lineSlice); i++ {
		if len(rigaper)+len(lineSlice[i]) < numero {
			rigaper += lineSlice[i] + " "
		} else {
			righe = append(righe, rigaper)
			rigaper = lineSlice[i] + " "
			continue
		}
	}
	return
}
