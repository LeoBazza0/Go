package main
import (
  "fmt"

)

func main () {
  var n int
  fmt.Print ("inserisci un numero: ")
  fmt.Scan (&n)
  n1:=1
  n2:=1
  n3:=0
  fmt.Print(n1, "+", n2)
  for i:=0; i<=n; i++ {
    n3=n1+n2
    n1=n2
    n2=n3
    fmt.Print("+", n3)
  }
  fmt.Print("\n", n3)
}
