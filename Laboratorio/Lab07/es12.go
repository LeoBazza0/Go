package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Leggere un testo (una sequenza di "parole", cioè di stringhe separate da spazi)
fino a incontrare la parola "stop", e stampare il testo senza punteggiatura
 (ciao, come stai? -> ciao come stai)
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
  //newLine := strings.Trim(line, "!.,<>=/|()[]@")
  sliceParole:= strings.Fields (line)
  for i :=0 ; i<len(sliceParole); i++ {
    fmt.Print(strings.TrimRight(sliceParole[i], "!:.,<>=/|()[]@"), " ") 
  }
  }
