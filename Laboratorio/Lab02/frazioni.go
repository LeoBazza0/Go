package main
import "fmt"
func main () {
  var n1, d1, n2, d2 int
  fmt.Print ("Inserie prima frazione: ")
  fmt.Scan (&n1, &d1)
  fmt.Print ("Inserie seconda frazione: ")
  fmt.Scan (&n2, &d2)
  if n1*d2==d1*n2 {
    fmt.Println ("equivalenti")
  } else {
    fmt.Println ("non equivalenti")
  }
}

//a:b=c:d
//a*d=b*c
