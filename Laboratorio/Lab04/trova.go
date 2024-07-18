package main
import "fmt"
func main () {
  var r rune
  var s string

fmt.Print ("Inserire una Stringa: ")
fmt.Scan (&s)
fmt.Print ("Inserie un carattere da trovare: ")
fmt.Scan (&r)

for i:=0; i<=len(s); i++ {
  if int(s[i])==int(r) {
  fmt.Print (s[i])
  break
} else {
  fmt.Print ("-1")
  break
}
  }
}
