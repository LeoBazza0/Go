package main
import "fmt"
func main () {
  var Celsius int
  const const1 = float64(9.0/5.0)
  fmt.Print ("Inserire gradi Celsius: ")
  fmt.Scan (&Celsius)
  fahrenheit:= (float64(Celsius)*const1)+32.0
  fmt.Println("in fahrenheit sono: ", fahrenheit)
  //fmt.Println (const1)
}
