package main
import (
  "fmt"
)

func main () {
  var n int
  maxSum:=0
  fmt.Print("inserisci un numero: ")
  for {
  if n==999 {
      break
    }
    maxSum2:=0
    for n>0 {
      digit:= n%10
      n /= 10
      maxSum2 += digit
    }
    if maxSum2>=maxSum {
        maxSum=maxSum2
    }
    fmt.Scan (&n)
  }
fmt.Println(maxSum)
}
