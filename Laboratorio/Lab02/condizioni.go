package main
import "fmt"
import "math"
func main () {
  var n int
  var m float64

  fmt.Print ("inserisci un valore uguale a 10: ")
  fmt.Scan (&n)
  fmt.Println (n==10)

  fmt.Print ("inserisci un valore diverso da 10: ")
  fmt.Scan (&n)
  fmt.Println ( n!=10)

  fmt.Println ("inserisci un valore diverso da 10 e da 20: ")
  fmt.Scan (&n)
  fmt.Println (n!=10 && n!=20)

  fmt.Print ("inserisci un valore diverso da 10 o da 20: ")
  fmt.Scan (&n)
  fmt.Println (n!=10 || n!=20)

  fmt.Print ("inserisci un valore maggiore o uguale a 10: ")
  fmt.Scan (&n)
  fmt.Println (n>=10)

  fmt.Print ("inserisci un valore compreso tra 10 e 20, nell'intervallo [10,20]: ")
  fmt.Scan (&n)
  fmt.Println(n>=10 && n<=20)

  fmt.Print ("inserisci un valore int compreso tra 10 e 20, nell'intervallo (10,20): ")
  fmt.Scan (&n)
  fmt.Println (n>10 && n <20)

  fmt.Print ("inserisci un valore compreso tra 10 e 20, nell'intervallo [10,20): ")
  fmt.Scan (&n)
  fmt.Println (n>=10 && n<20)

  fmt.Print ("inserisci un valore esterno all'intervallo [10,20]: ")
  fmt.Scan (&n)
  fmt.Println (n<10 && n>20)

  fmt.Print ("inserisci tra 10 e 20 (nell'intervallo [10,20]) o tra 30 e 40 ([30,40]): ")
  fmt.Scan (&n)
  fmt.Println ( (n>=10 && n<=20) || (n>=30 && n<=40) )

  fmt.Print ("inserisci un valore multiplo di 4 ma non di 100: ")
  fmt.Scan (&n)
  fmt.Println (n%4==0 && !(n%100==0))

  fmt.Print ("inserisci un valore dispari e compreso tra 0 e 100 ([0,100]): ")
  fmt.Scan (&n)
  fmt.Println ((n>=0 && n<=100) && (n%2==0))

  fmt.Print ("inserisci un valore float64 vicino a 10.0 (non più lontano di 10^-6) (1e-6): ")
  fmt.Scan (&m)
  fmt.Println (math.Abs(m-10.0)<=1e-6)

}
