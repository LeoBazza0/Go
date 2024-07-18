package main
import "fmt"
func main () {
var n int
fmt.Print ("Inserire numero codice: ")
fmt.Scan (&n)
if n==1 {
  fmt.Println ("caricamento acqua")
} else if n==2 {
  fmt.Println ("scarico acqua")
} else if n==3 {
  fmt.Println ("sistema di riscaldamento")
} else if n<=0 || n>3 {
  fmt.Println ("errore")
}
}
