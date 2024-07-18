package main
import "fmt"
func main () {
  var o, m int
  fmt.Print ("Insterire un orario: ")
  fmt.Scan (&o, &m)
  if m>=30 && o<=12 {
    fmt.Println ("mattino")
  }
  if m<=30 && o>=5 {
    fmt.Println ("mattino")
  }
}
