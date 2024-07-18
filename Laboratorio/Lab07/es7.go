package main
import (
  "fmt"
  "bufio"
  "os"
  "sort"
  "strings"
)

/*
Leggere un testo (una sequenza di "parole", cioè di stringhe separate da spazi)
e stampare le parole in ordine alfabetico.
*/

func main () {
  scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

 parole := strings.Fields (line)
 if parole[i]=="stop"{
   break
 }
fmt.Println (parole)
sort.Strings (parole)
fmt.Println (parole)
}
