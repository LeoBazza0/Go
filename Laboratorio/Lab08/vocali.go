package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stringa string
var vocali map[rune]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringa = scanner.Text()
	vocali = make(map[rune]int)
	contaVocali(stringa, vocali)
}

func contaVocali(stringa string, vocali map[rune]int) {
	for _, r := range "aeiouAEIOU" {
		vocali[r] = strings.Count(stringa, string(r))
	}
	for k, v := range vocali {
		fmt.Println(string(k), ":", v)
	}
}
