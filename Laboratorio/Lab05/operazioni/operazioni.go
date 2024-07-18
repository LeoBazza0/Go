/*

In un file operaazioni.go definire le seguenti funzioni:

- operazioni(n1, n2 int) (int, int, int)
che accetta due interi e restituisce somma, prodotto e differenza.
Scrivine una versione con variabili di ritorno nominate e una senza (puoi commentare una delle due versioni per caricare un file solo)

Scrivi un main per invocare e testare la funzione. Il programma legge da standard input due int.

nomefile: operazioni.go

*/


package main
import (
  "fmt"
)

func main () {
  var n1, n2 int
  fmt.Println ("inserisci due numeri")
  fmt.Scan (&n1, &n2)
  fmt.Println (operazioni(n1, n2))
}

//funzione con variabili di ritorno non nominate
func operazioni(n1, n2 int) (int, int, int) {
  return n1+n2, n1*n2, n1-n2
}

/*
//funzione con variabili di ritorno nominate
func operazioni(n1, n2 int) (somma int, prodotto int, differenza int) {
  return n1+n2, n1*n2, n1/n2
}

*/
