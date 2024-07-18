package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const N = 10
	parole := make([]string, N)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parole = strings.Split(line, " ")
		fmt.Println("line: ", line)
	}
	fmt.Println(parole)
	piuCorta(parole)
}

func piuCorta(parole []string) {
	var lunghezzaMax, iMax int
	for i := 0; i < len(parole); i++ {
		if len([]byte(parole[i])) > lunghezzaMax {
			lunghezzaMax = len([]byte(parole[i]))
			iMax = i
			fmt.Println(i, parole[i], len([]byte(parole[i])))
			continue
		}
	}
	fmt.Println("la parola più corta in temrini di byte è: ", parole[iMax])

}

//sbagliato aaa
