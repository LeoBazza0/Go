/*

package main
import (
  "fmt"

)

func main () {
//var
  var stringaS string
  var stringaAR []rune
  var old, new rune
//istruioni
  fmt.Print ("inserire la stringa: ")
  fmt.Scan(&stringaS)
  fmt.Print ("inserire la vecchia runa: ")
  fmt.Scan (&old)
  fmt.Print ("inserire la nuova runa: ")
  fmt.Scan (&new)
for _, k:= range stringaS {
  stringaAR= append (stringaAR, rune(k))
}
//funzione repInPlace
  repInPlace(&stringaAR, old, new)
  output:=string(stringaAR)
  fmt.Println("stringa modificata:", output)
}

func repInPlace(stringa *[]rune, old rune, new rune) {
	for i, r := range *stringa {
		if r == old {
			(*stringa)[i] = new
		}
	}
	return
}
*/


package main

import (
	"fmt"
	"os"
)

func repInPlace(stringa *[]rune, old rune, new rune) {
	for i, r := range *stringa {
		if r == old {
			(*stringa)[i] = new
		}
	}
	return
}

func main() {
	strToModify := []rune(os.Args[1])
	old := []rune(os.Args[2])[0]
	new := []rune(os.Args[3])[0]
	fmt.Println(os.Args[1], string(old), string(new))
	repInPlace(&strToModify, old, new)
	output := string(strToModify)
	fmt.Println("stringa modificata:", output)
}
