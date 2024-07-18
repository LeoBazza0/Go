package main
import "fmt"
func main () {
  var p, a float64
  fmt.Print ("Inserisci peso (in kg): ")
  fmt.Scan (&p)
  fmt.Print ("Inserisci altezza (in cm): ")
  fmt.Scan (&a)
  if (p>=10 && p<50) && (a>=100 &&a<150) {
    fmt.Println ("normopeso")
  } else if (p>=10 && p<50) && (a>=150 &&a<=200) {
    fmt.Println ("sottopeso")
  } else if (p>=50 && p<=100) && (a>=100 &&a<=150) {
    fmt.Println ("sovrappeso")
  } else if (p>=50 && p<=100) && (a>150 &&a<=200) {
    fmt.Println ("normopeso")
  } else if p>100 && (a>=100 &&a<=150) {
    fmt.Println ("molto sovrappeso")
  } else if p>100 && (a>150 &&a<=200) {
    fmt.Println ("sovrappeso")
  }
}
