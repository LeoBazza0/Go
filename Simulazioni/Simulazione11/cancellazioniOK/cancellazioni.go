package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cancella(lista []string) (newList []string) {
	for i := 0; i < len(lista); i += 0 {
		valore, err := strconv.Atoi(lista[i])
		if err == nil {
			i += valore + 1
		} else {
			newList = append(newList, lista[i])
			i++
		}
	}
	return
}

func main() {
	var line string
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		os.Exit(0)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File non accessibile")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line += scanner.Text() + " "
	}
	lista := strings.Fields(line)
	fmt.Println(lista)
	fmt.Println(cancella(lista))
}
