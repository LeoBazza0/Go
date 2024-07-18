package main
import "fmt"
func main () {
  var o int
  fmt.Print ("Insterire un orario: ")
  fmt.Scan (&o)
  if o>=6 && o<13 {
    fmt.Println ("mattino")
  }
}
