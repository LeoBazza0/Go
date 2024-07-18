package main
import "fmt"

/*
//con variabili non nominate
func operazioni(n1, n2 int) (int, int, int) {
  somma:=n1+n2
  prodotto:=n1*n2
  differenza:=n1-n2
  return somma, prodotto, differenza
}
*/

//con variabili nominate
func operazioni(n1, n2 int) (somma int, prodotto int, differenza int) {
  somma=n1+n2
  prodotto=n1*n2
  differenza=n1-n2
  return somma, prodotto, differenza
}


func main () {
var n1, n2 int
fmt.Print ("inserire due numeri: ")
fmt.Scan (&n1, &n2)
fmt.Println(operazioni(n1, n2))
}
