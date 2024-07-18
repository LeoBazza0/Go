package main
import (
  "fmt"

)

const DIM = 5


func main () {
  const DIM = 5
  var array [DIM]int
  for i:=0 ; i<DIM ; i++ {
    array[i]= i
  }
  fmt.Println ("array: ", array)
  reverse(&array)
  fmt.Println ("reverse: ", array)
  switchFirstLast(&array)
  fmt.Println ("switchFirstLast: ", array)

}

func reverse(a *[DIM]int) {
	len := len(a)
	for i := 0; i < len/2; i++ {
		a[i], a[len - 1 - i] = a[len - 1 - i], a[i]
	}
}

func switchFirstLast (a *[DIM]int) {
  ultimo:=a[DIM-1]
  primo:=a[0]
  a[DIM-1]=primo
  a[0]=ultimo
  }
