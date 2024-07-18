package main
import "fmt"
func main () {
  var r, p string
  fmt.Print ("Inserisci 5 caratteri ")
  // var n byte
  for i:=0; i<5; i++ {
    fmt.Scan (&r)
    if r>p {
      p=r
      }
  }
  fmt.Print (p)
}
