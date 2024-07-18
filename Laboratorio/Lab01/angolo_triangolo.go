package main
import "fmt"
func main () {
  var a1, a2, a3 float64
  const tot int = 180
  // fmt.Println ("Inserisci angolo 1")
  fmt.Scan (&a1)
  // fmt.Println ("Inserisci angolo 2")
  fmt.Scan (&a2)
  a3=float64(tot)-((a1+a2))
  fmt.Println ("Angoli uno e due: ", a1, a2, )
  fmt.Println ( "Ampiezza terzo angolo: ", a3)
}
