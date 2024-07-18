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
 e stampare solo le parole di lunghezza maggiore della media delle lunghezze di tutte le parole del testo.
*/

func main () {
  var media, mediaSomma, j int
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  line := scanner.Text()
  paroleSlice:=strings.Fields (line)
  for i:=0; i<len(paroleSlice); i++ {
    mediaSomma=mediaSomma+len(paroleSlice[i])
    j++
  }
  media=mediaSomma/j
  fmt.Println(media)
  for i:=0; i<len(paroleSlice); i++ {
    if len(paroleSlice[i])>=media {
      fmt.Println (paroleSlice[i])
    }
}
}
