package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//deve leggere da standard input

func main() {
	var max int = 0
	var e int
	var stringa_iniziale []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numeri := scanner.Text()
	stringa_iniziale = strings.Split(numeri, " ")

	for i := 0; i < len(stringa_iniziale); i++ {
		if len(stringa_iniziale[i]) > max {
			max = len(stringa_iniziale[i])
			e = i
		} else if len(stringa_iniziale[i]) == max {
			e = i
		}
	}
	fmt.Println(stringa_iniziale[e])
}

//1. devono essere forniti da standard imput
//2. i dati forniti sono di tipo string
//3. i dati forniti sono in serie e sono inseriti in una slice
//4. i dati forniti sono solo da inserire in una slice ma per il resto sono pronti da usare
//5. è indispensabile avere in memoria la lunghezza della parola più lunga
//6. il programma deve fermarsi una volta premuto invio, quindi legge solo la prima riga, se avessimo
//	voluto far leggere anche le altre righe avremo dovuto mettere un "for" nella fase di lettura
