package main
import "fmt"
func main () {
  var val1, val2 intk
  var _, err = fmt.Scan(&val1, &val2)
  fmt.Println ("val1 e val2: ", val1, val2)
  fmt.Println ("errore: ", err)
  fmt.Println ("prima dello scambio: ", val1, val2)
  val1, val2 = val2, val1
  fmt.Println ("dopo lo scambio: ", val1, val2)
}
