package main
import "fmt"
func main () {
  var max, min int
  fmt.Print ("Inserire due numeri: ")
  fmt.Scan (&max, &min)
  if min>max {
    min, max=max, min
  }
  fmt.Println (max)
}
