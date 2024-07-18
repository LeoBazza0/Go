
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 0 nome, 1 (sigla), 2 regione, (popolazione), (superficie), (densita), (altitudine)
 type capoluoghi struct {
	nome, regione string
}

var m map[string]string
var text, nome, regione, regioneimput, nome1 string

func main() {
	// var cap capoluoghi
	var lineslice, nomislice []string
	m = make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToUpper(line)
		lineslice = strings.Split(line, ",")
		for i := 0; i < 8; i++ {
			regione = lineslice[2]
			nome = lineslice [0]
			if regione==regioneimput {
				nomislice=append(nomislice, nome)
				fmt.Print (regione, nome)
			}
		}

	}
	for j := 1; j < len(os.Args); j++ {
		regioneimput = os.Args[j]
	fmt.Println (m[regione])
	}
	fmt.Println (nomislice)

}

*/

/////////////////////////////////////////////////////////////

package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Capoluogo struct {
    nome, regione string
}

const POS_NOME = 0
const POS_REGIONE = 2
func main() {
    var regione string
    var capoluoghi map[string]([]Capoluogo)
    capoluoghi = make(map[string]([]Capoluogo))
		scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        line := scanner.Text()
        campi := strings.Split(line, ",")
				capoluoghi[campi[POS_REGIONE]] = append(capoluoghi[campi[POS_REGIONE]], Capoluogo{nome: campi[POS_NOME], regione: campi[POS_REGIONE]})
}
for i := 1; i < len(os.Args); i++ {
			 regione = os.Args[i]
			 listaCapoluoghiInRegione, ok := capoluoghi[regione]
			 if !ok {
					 fmt.Println("Input", regione, "non valido")
					 continue
			 }
			 fmt.Printf("Capoluoghi della regione %s:\n", regione)
        for _, cap := range listaCapoluoghiInRegione {
            fmt.Printf("%s ", cap.nome)
        }a
    }
}
