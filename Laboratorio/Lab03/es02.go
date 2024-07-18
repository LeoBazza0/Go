package main
import "fmt"
func main () {
  const k=5
  var n int
  for i:=1; i<=k; i++ {
    fmt.Print ("Inserire un numero: ")
    fmt.Scan (&n)
    fmt.Println (2*n)
  }
}
