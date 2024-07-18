/**

Esercizio statistica
====================

Scrivere un programma in Go (il file deve chiamarsi 'statistica.go') che,
data una serie (di lunghezza arbitraria) di numeri interi da standard input
(la serie contiene almeno un elemento, e termina con CTRL-D su una riga vuota),
calcola ed emette su standard output:

- il valore minimo (int)

- il valore massimo (int)

- la media (in tipo float64, quindi non troncata)

dei valori letti. Per il formato si vedano gli esempi qui sotto.

Si fornisce un file "esempioLungo.input" d'esempio, a fronte del quale l'output *dovrà*
essere il seguente:

min: -3556
max: 9292
media: 2871.39393939394

E un file "esempioBreve.input" d'esempio, a fronte del quale l'output *dovrà*
essere il seguente:

min: -26
max: 77
media: 27.75

*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestInput(t *testing.T) {
	lancia("esempioLungo.input", "min: -3556\nmax: 9292\nmedia: 2871.3939393939395\n", t)
	lancia("esempioBreve.input", "min: -26\nmax: 77\nmedia: 27.75\n", t)
	lancia("esempioPos.input", "min: 11\nmax: 77\nmedia: 35.42857142857143\n", t)
	lancia("esempioNeg.input", "min: -77\nmax: -11\nmedia: -35.42857142857143\n", t)
}

func lancia(input, expOut string, t *testing.T) {
	//fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./statistica") // presuppone che sia già stato compilato (dallo script)

	in, err := os.Open(input)
	subproc.Stdin = in
	out, err := subproc.CombinedOutput()

	if err != nil {
		// verificare se uscito per esplicito os.Exit()
		//t.Errorf("Fallito: %s\n", stderr)
		fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se l'esercizio richiedeva uscita esplicita con exit code)\n", err)
	}

	fmt.Printf("Output:\n%s\n", string(out))
	fmt.Printf("Expected output:\n%s\n", string(expOut))

	if string(out) != string(expOut) {
		fmt.Println(">>> FAIL!")
		t.Fail()
	} else {
		fmt.Println(">>> PASS!")
	}
	fmt.Println("-----")

	subproc.Process.Kill()
	//subproc.Wait()
}
