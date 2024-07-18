package main
import (
  "fmt"
)

var base, altezza int
type rettangolo struct {
  base,altezza int
}

func main () {
  var rett rettangolo
//  var err error
  fmt.Print ("Inserisci base rettangolo: ")
  fmt.Scan (&rett.base)
  fmt.Print ("Inserisci altezza rettangolo: ")
  fmt.Scan (&rett.altezza)

  disegno:=rett.String()
fmt.Print(disegno)
}

func (rett rettangolo) String() string {
  var disegno string
  for i:=0; i<rett.altezza; i++ {
    for j:=0; j<rett.base; j++ {
      disegno+="."
    }
    disegno+="\n"
  }
  return disegno
}
