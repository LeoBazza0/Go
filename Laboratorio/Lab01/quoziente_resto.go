package main
import "fmt"
func main () {
  var n1, n2 int
  fmt.Print ("Inserire dividendo: ")
  fmt.Scan (&n1)
  fmt.Print ("Inserire divisore: ")
  fmt.Scan (&n2)
    q:=n1/n2
    r:=n1%n2
  fmt.Println ("Quoziente: ", q)
  fmt.Println ("Resto: ", r)
}
