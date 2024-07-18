package main
import "fmt"
func main () {
  var n int
  fmt.Print ("Inserie orario: ")
  fmt.Scan (&n)
  if n<0 || n>=24 {
    fmt.Println ("Orario non valido")
  } else if n>=0 && n<=6 {
    fmt.Println ("Notte")
  } else if n>6 && n<=13 {
    fmt.Println ("Mattina")
  } else if n>13 && n<=18 {
    fmt.Println ("Pomeriggio")
  } else if n>18 && n<=23 {
    fmt.Println ("Sera")
  }
}
