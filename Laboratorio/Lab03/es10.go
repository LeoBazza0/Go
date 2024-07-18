package main
import (
  "fmt"
"math/rand"
"time"
)
 func main () {
   var n, j int
   const TARGET=100

   rand.Seed(time.Now().UnixNano())
   for j<=TARGET {
     fmt.Print(n, " ")
       n=rand.Intn(11)
     j= j+n
   }
   fmt.Print (j)
 }
