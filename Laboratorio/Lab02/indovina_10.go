package main
import "fmt"
func main () {
  var n int
//  numero = 8
  fmt.Print ("Inserire il numero: ")
  fmt.Scan (&n)
  if n<1 || n>10 {
    fmt.Println ("Valore non valido")
  } else if n!=8  {
    fmt.Println ("Non hai indovinato")
  } else if n==8 {
    fmt.Println ("Hai indovinato")
  }
}

// if n==8
