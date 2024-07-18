package main
import "fmt"
func main () {
  var x, y, m, q float64
  fmt.Print ("Inserisci valore X del punto: ")
  fmt.Scan (&x)
  fmt.Print ("Inserisci valore Y del punto: ")
  fmt.Scan (&y)
  fmt.Print ("Inserisci valore M della retta: ")
  fmt.Scan (&m)
  fmt.Print ("Inserisci valore Q della retta: ")
  fmt.Scan (&q)
  // y=mx+q
  if y==m*x+q {
    fmt.Println ("Sulla retta")
  } else if y<m*x+q {
    fmt.Println ("Sotto")
  } else if y>m*x+q {
    fmt.Println ("Sopra")
  } else {
  }
}
