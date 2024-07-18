/*
Scrivere un programma stampaAlternata.go che legge da standard input del testo su più righe
 (terminato da EOF) e stampa prima le righe pari e poi le righe dispari
 (considerate la prima riga del testo la riga 1 (e non 0)).
*/

package main
import (
  "fmt"
  "bufio"
  "os"
)

func main () {
  lineeDispari := make([]string, 0)
	lineeLette := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineeLette++
		if lineeLette%2 == 0 {
			fmt.Println(scanner.Text())
		} else {
			lineeDispari = append(lineeDispari, scanner.Text())
		}
	}
	for _, line := range lineeDispari {
		fmt.Println(line)
	}
}
