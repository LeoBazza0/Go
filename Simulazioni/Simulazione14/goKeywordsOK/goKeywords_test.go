/**

Esercizio "parole riservate"
===========================

Scrivere un programma Go (il file deve chiamarsi 'goKeywords.go') che identifica le parole riservate
in un programma Go, e per ciascuna stampa le linee di codice in cui si trova.
Se una parola chiave non compare nel testo non viene stampato nulla.
La stampa delle occorrenze delle parole chiave deve seguire l'ordine alfabetico delle stesse.

In particolare il programma deve essere dotato di:

- un funzione printKeywords
che stampa, una per riga e in ordine alfabetico, la lista delle parole riservate trovate, ciascuna seguita dai numeri di linea di codice in cui è stata trovata

- la funzione main
che legge da un file, il cui nome è passato da linea di comando, il testo (un programma Go) da analizzare.
Se manca il parametro da linea di comando, il programma stampa "A file name, please"
Se il file non esiste o non viene aperto correttamente, il programma stampa "File not found".
Altrimenti il programma stampa, una per riga e in ordine alfabetico, la lista delle parole riservate trovate nel file, ciascuna seguita da ':'
e dalla lista dei numeri di linea di testo (del programma) in cui è stata trovata.

Elenco delle keywords di Go

"break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for",
"func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select",
"struct", "switch", "type", "var"

Esempi
------

$ ./goKeywords hello_world.go.sample
func: [5 9]
import: [3]
package: [1]

dove hello_world.go.sample è il programma che trovate in questa dir.



$ ./goKeywords if.go.sample
else: [9 19 21]
func: [5]
if: [7 13 17 19]
import: [3]
package: [1]

dove if.go.sample è il programma che trovate in questa dir.

**/

package main

import (
	//	"fmt"
	//	"math"
	"testing"
)

var prog = "./goKeywords"

func TestComeDaEsempio1(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "func: [5 9]\nimport: [3]\npackage: [1]\n", "hello_world.go.sample")
}

func TestComeDaEsempio2(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "else: [9 19 21]\nfunc: [5]\nif: [7 13 17 19]\nimport: [3]\npackage: [1]\n", "if.go.sample")
}

/*func TestEsistenzaIsKeywords(t *testing.T) {
	// fallisce compilazione se manca fn
	isKeyword("break")
}*/

func TestEsistenzaPrintKeywords(t *testing.T) {
	tmp := make(map[string][]int)
	tmp["word"] = append(tmp["word"], 10)
	// fallisce compilazione se manca fn
	printKeywords(tmp)
}

//func TestMancaArgomento DA IMPLEMENTARE

//func TestMancaFile DA IMPLEMENTARE
