package main
import "fmt"

func main () {
  var s string
  var n int
  const topLen = 10

  fmt.Print ("Inserisci stringhe fino al raggiungimento del limite: ")
  for {
    fmt.Scan (&s)
    n+=len(s)
    if n>=topLen {
      break
    }
  }
}
