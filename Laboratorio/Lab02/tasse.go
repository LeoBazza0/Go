package main
import "fmt"
func main () {
  var r int
  var c bool
  const (
  NC=32000
  SC=64000
  P1=10
  P2=24
  )
  fmt.Print ("Inserire reddito: ")
  fmt.Scan (&r)
  fmt.Print ("Inserire se coniugato: ")
  fmt.Scan (&c)
  if r<=NC && c==false || r<=SC && c==true {
    fmt.Println ("Percentuali tasse da pagare: ", P1, "%")
  } else if r>NC && c==false || r>SC && c==true {
    fmt.Println ("Percentuali tasse da pagare: ", P2, "%")
    }
}
