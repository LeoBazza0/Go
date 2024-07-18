package main
import "fmt"
import "math"
func main () {
  var x, y float64
  fmt.Println ("Inserire i valori di x e y: ")
  fmt.Scan (&x, &y)
  const  EPSILON = 1e-5
  if math.Sqrt(x*x+y*y)<=EPSILON {
    fmt.Println ("vicino all'origine")
  } else {
    fmt.Println ("non vicino all'origine")
  }
}
