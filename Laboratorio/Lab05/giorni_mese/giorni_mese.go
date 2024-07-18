package main
import (
  "fmt"
  "strconv"
  "strings"
)

func main () {
  var s string
  fmt.Print("inserisci un mese (gg/mm/aaaa): ")
  fmt.Scan (&s)
  sliceData:=strings.Split(s,"/")
  mese:=sliceData[1]
  meseInt, _ := strconv.Atoi(mese)
  fmt.Println ("il mese ", meseInt, " ha ", giorniInMese(meseInt), "giorni")
}

//s[strings.Index(s, "/") + 1 : strings.LastIndex(s, "/")] --> "mm"  s [3,5)

func giorniInMese (mese int) int {
  switch mese {
    case 2: return 28
    case 4, 6, 9, 11: return 30
    default: return 31
  }
}



/*
Vogliamo scrivere una funzione
	giorniInMese(mese int) int
che, dato come parametro il numero corrispondente a un mese, restituisce il numero di giorni di quel mese
 (28 per febbraio; 30 per aprile, giugno, settembre, novembre; 31 per G,M,M,L,A,O,D).
 Assumiamo che il numero passato come parametro sia in [1,12], quindi non facciamo controlli.

 Scrivi la funzione giorniInMese come da specifiche. Scrivi un main per invocare e testare la funzione.
 Il programma deve leggere da standard input una stringa nel formato gg-mm-aaaa (vedi funzione Atoi del pacchetto strconv)
 e stampare "il mese <x> ha <tot> giorni", dove x e tot sono numeri, usando la funzione giorniInMese per determinare tot.
  Chiama il programma giorni_mese.go e caricalo su upload
  */
