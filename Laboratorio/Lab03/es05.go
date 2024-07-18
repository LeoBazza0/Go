package main
import "fmt"
func main() {
  var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
  for i:=1; i<=n; i++ {
    fmt.Print(i, " ")
  }
}
