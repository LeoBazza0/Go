package main
import "fmt"
func main () {
  var s string
  var spv byte
  fmt.Print("Inserie stringa: ")
  fmt.Scan (&s)
  for i:=0; i<len(s); i++ {
    fmt.Print (string(s[i]))
   if s[i]<spv {
    fmt.Print("-")
    }
    spv=s[i]
  }
}
