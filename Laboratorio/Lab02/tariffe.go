package main
import "fmt"
func main () {
  var e int
  var s bool
fmt.Print ("inserie l'età: ")
fmt.Scan (&e)
fmt.Print ("Studente? ")
fmt.Scan (&s)
//  if s==false {
    if e>=0 && e<9 {
      tariffa:="ingresso gratis"
      fmt.Println (tariffa)
      } else if e>=9 && e<18 || s==true && e>=18 {
        tariffa:= "ingresso 5 euro"
        fmt.Println (tariffa)
      } else if (e>=18 && e<26) || e>=65 {
          tariffa:="ingresso 7.5"
          fmt.Println (tariffa)
      } else if e>=26 && e<65 {
        tariffa:="ingresso 10"
        fmt.Println (tariffa)
      }
//  }
}
