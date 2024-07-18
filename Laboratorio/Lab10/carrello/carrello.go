package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//strings.Split(imput, "|") []string
//num, err:=strconv.Atoi([]string[i])

type carrello struct {
	pos    int
	carico int
}

const pesoMax = 15

func main() {
//varuabili
	var trackInt []int
	var stringInput string
	var count int
//uso le funzioni che ho creato
	fmt.Scan(&stringInput)
	trackInt = readInput(stringInput) //creo una slice di Int
	printSlice(trackInt) //stampo la slice di Int

	for i:=0; i<len(trackInt); i++ {

		if count>pesoMax {
			trackInt[i]
			printSlice(trackInt)

		}
		count+=trackInt[i]
		trackInt[i]=0

	}
}

func readInput(string) (trackInt []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = line[0 : len(line)-1]
	//elementi è una slice di stringhe con un numero per cella
	elementi := strings.Split(line, "|")
	fmt.Println(elementi)
	for _, value := range elementi {
		if value == " " {
			trackInt = append(trackInt, 0)
		} else {
			num, _ := strconv.Atoi(value)
			trackInt = append(trackInt, num)
		}
	}
	return trackInt
}
func printSlice(trackInt []int) {
	for i := 0; i < len(trackInt); i++ {
		fmt.Print("|")
		fmt.Print(trackInt[i])
	}
	fmt.Print("| \n")
}
func (c Carrello) String() string {
    return fmt.Sprintf("posizione %d, carico %d", c.pos, c.carico)
}



/*
func aggiornaStato(c *Carrello, posizione, carico int) bool {

}
*/

/*
func update the track (in un altro modo ma che non funziona)
var count, pos int
for pos, _ = range trackInt {
	if count < pesoMax {
		count += trackInt[pos]
		trackInt[pos] = 0
	} else {
		trackInt[0] += count
		count = 0
	}
	printSlice(trackInt)
/*
leggo input
ciclo della slice, con un contatore++, eliminando i caratteri analizzati
se contatore andrebbe sopra il pesoMax mi fermo, aggiungo alla casella 0 il carico ed elimino i caratteri analizzati
ogni volta ristampo la cosa
*/
