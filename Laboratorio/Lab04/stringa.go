package main
import "fmt"
 func main () {
   var s string
   var spv byte
fmt.Print("Inserie stringa: ")
fmt.Scan (&s)
   for i:=0; i<len(s); i++ {
    if s[i]<spv {
     fmt.Print("-")
    } else if s[i]>spv {
      fmt.Print("+")
    } else {
      fmt.Print("=")
    }
    spv=s[i]
   }
 }
