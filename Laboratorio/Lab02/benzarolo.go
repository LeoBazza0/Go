package main
import "fmt"
func main () {
  var l float64
  var t int
  t0:=1.78
  t1:=1.98
  t2:=1.2
  t3:=1.1
  fmt.Print ("Inserire quanti litri: ")
  fmt.Scan (&l)
  fmt.Print ("Inserire il tipo di carburante: ")
  fmt.Scan (&t)
  if t==0 {
    f:=l*t0
    fmt.Print ("Il costo finale è: ", f)
  } else if t==1 {
    f:=l*t1
    fmt.Print ("Il costo finale è: ", f)
  } else if t==2 {
    f:=l*t2
    fmt.Print ("Il costo finale è: ", f)
  } else if t==3 {
    f:=l*t3
    fmt.Print ("Il costo finale è: ", f)
  } else {
    fmt.Print ("Il tipo di carburante inserito è errato")
  }
}
