package main
import (
"fmt"
"math/rand"
"time"
)
func main () {
  var n int
  const K=10
  for i:=1; i<=K; i++ {
  //  rand.NewSource(time.Now().UnixNano())
     rand.Seed(time.Now().UnixNano())
    n=rand.Intn(11)
    fmt.Print (n, " ")
  }
}
