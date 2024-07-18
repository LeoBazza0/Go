package main
import (
  "fmt"

)

type segmento struct {
  estremi byte
  interno byte
  orizzontale bool
  lunghezza int
}

func main () {
  var seg segmento
  var estremi, interno string
  var orizzontale bool
  var lunghezza int
  fmt.Print ("Inserisci simboli per estremi: ")
  fmt.Scan (&estremi)
  fmt.Print("Inserisci i simboli per interni: ")
  fmt.Scan (&interno)
  fmt.Print ("Orizzontale (True) o verticale (False)? ")
  fmt.Scan (&orizzontale)
  fmt.Print ("Inserisci lunghezza segmento: ")
  fmt.Scan (&lunghezza)


  seg = segmento{estremi[0],interno[0],orizzontale,lunghezza}
  disegno:=seg.String()
  fmt.Println(disegno)
}

//ora faccio qualcosa per farlo verticale
//strings.repeat
func (seg segmento) String() string {
  var disegno string
  disegno+=string(seg.estremi)
  for i:=2; i<seg.lunghezza; i++ {
    disegno+=string(seg.interno)
  }
  disegno+=string(seg.estremi)
  return disegno
}
