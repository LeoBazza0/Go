package main
import "fmt"
func main () {
  var n int
  fmt.Print ("Inserire un numero per disegnare la V ")
  fmt.Scan (&n)
for i:=0; i<=n; i++ {
  fmt.Print("*")
  for j:=0; j<n-2; j++ {
    fmt.Print (" ")
  }
}


  }
