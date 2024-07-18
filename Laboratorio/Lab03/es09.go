package main
import "fmt"
func main() {
  var n,j int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
  for i:=1; j<n; i++ {
    fmt.Println (j)
    j=i*i
    }

}
