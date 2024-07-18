package main
import (
  "fmt"
  "os"
)

func main () {
  var count int
  parola:=os.Args[1]
  carattere:=os.Args[2]
  for _,k :=range parola {
    if string(k)==carattere{
      count++
    }
  }
  fmt.Println (count)
}
