package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

/*
Leggere un testo (una sequenza di "parole", cioè di stringhe separate da spazi)
fino a incontrare EOF (ctrl D su una riga nuova)
e stampare solo (e tutte) le parole che hanno lunghezza massima.
*/

func main () {
  var lunghezzaMaxSlice int
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  line := scanner.Text() //string lunga e con spazi
  paroleSlice:=strings.Fields (line) // slice con una parola per "posto"
  for i:=0; i<len(paroleSlice); i++ {
    if len(paroleSlice[i])>lunghezzaMaxSlice {
      lunghezzaMaxSlice=len(paroleSlice[i])
    }
  }
  for j:=0; j<len(paroleSlice); j++ {
    if len(paroleSlice[j])==lunghezzaMaxSlice {
      fmt.Println (paroleSlice[j])
    }
  }
}
