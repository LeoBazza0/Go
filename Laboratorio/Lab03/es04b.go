package main
import (
"fmt"
"math/rand"
"time"
)
func main () {
  var n, j int
  const K=10
  rand.Seed(time.Now().UnixNano())
  for i:=1; i<=K; i++ {
    n=rand.Intn(11)
    fmt.Print (n, " ")
    if n%2==0 {
      j++
    }
  }
  fmt.Print ("Ci sono ", j, " numeri pari")
}
