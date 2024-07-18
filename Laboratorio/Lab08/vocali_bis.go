package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "sort"
)

var stringa string
var vocali map[rune]int
var keys []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringa = scanner.Text()
	vocali = make(map[rune]int)
	contaVocali(stringa, vocali)

  for k:= range vocali {
    keys= append (keys, string(k))
  }
  sort.Strings (keys)
  for _, r:= range keys {
    runa:= []rune (r)
    fmt.Println (r, vocali [runa[0]])
  }
}

func contaVocali(stringa string, vocali map[rune]int) {
	for _, r := range "aeiouAEIOU" {
		vocali[r] = strings.Count(stringa, string(r))
	}
}
