package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
/*
Leggere un testo (una sequenza di "parole",
cioè di tringhe separate da spazi)
e stampare una parola per riga, partendo dall'ultima.
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
  parole:=strings.Fields (line)
  for i:=len(parole)-1; i>=0; i -- {
    if parole[i]=="stop"{
      break
    }
    fmt.Println (parole[i])
  }
  //fmt.Printf("%q \n", strings.Split(line, " "))
	// slice1:= []string {strings.Split (os.Args)}
}
