package main
import "fmt"
func main () {
  var h, m int
  for {
    fmt.Print ("Inserire ore: ")
    fmt.Scan (&h)
    fmt.Print ("Inserire minuti: ")
    fmt.Scan (&m)
    if (h>=0 && h<=23) && (m>=0 && m<=59) {
      fmt.Print ("valori validi")
      break
    }
  }
}
