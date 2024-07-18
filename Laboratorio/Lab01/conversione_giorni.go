package main
import "fmt"
func main () {
  //covertire un numero di gionri in anni-mesi-gionri
  var g int
  fmt.Print ("Inserire il numero di giorni da convertire: ")
  fmt.Scan (&g)
  const anni = 365
  const mesi = 30
//  const giorni = 7
  af:=g/anni
  mf:=(g%anni)/mesi
  gf:=(g%anni)%mesi
  fmt.Println ("anni", af, "mesi", mf,  "giorni", gf)
}
