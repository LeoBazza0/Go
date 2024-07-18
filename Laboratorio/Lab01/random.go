package main
import "fmt"
import "math/rand"
import "time"
func main () {
  var numero int
  rand.Seed(time.Now().Unix())
  numero=rand.Intn(101)
  fmt.Println (numero)
}
