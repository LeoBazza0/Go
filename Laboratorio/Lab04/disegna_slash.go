package main
import "fmt"
func main () {
  var n int
  fmt.Print ("Inserire un valore: ")
  fmt.Scan (&n)
  for i:=0; i<n; i++ {
    for j:=0; j<i; j++ {
      fmt.Print (" ")
    }
    fmt.Print ("*")
    fmt.Println ()
  }
}