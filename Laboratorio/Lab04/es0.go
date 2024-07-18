package main
import "fmt"

func main () {
  var n rune
  fmt.Print ("Inserie un carattere: ")
  fmt.Scanf ("%c", &n)
  fmt.Println (n, string(n))



  n2:=n-1
  n3:=n+1
  fmt.Println ("Byte precedente: ", /* n2, */ string(n2))
  fmt.Println ("Byte successivo: ", /* n3, */ string(n3))
/*
  if n2>=65 && n2<=76 {
    fmt.Println ("A-L")
  } else {
    fmt.Println ("altro")
  }

var s string
fmt.Scan (&s)
for i:=0; i<len(s); i++ {
fmt.Println (string(s[i]))
}
*/
}
