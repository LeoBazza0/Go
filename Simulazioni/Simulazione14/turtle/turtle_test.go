/**
Esercizio "tartaruga"
=====================

Scrivere un programma Go (il file deve chiamarsi 'turtle.go') che fa muovere una tartaruga nel piano cartesiano. La tartatura è comandata attraverso semplici
comandi (vedi sotto) e può solo fare passi unitari (interi) nelle 4 direzioni cardinali: Nord, Est, Sud, Ovest (alto, destra, basso, sinistra).

Il programma deve essere dotato di:

- una struttura Turtle con due campi x, y (coordinate nel piano cartesiano) di tipo int e un campo dir (direzione N, E, S, O) di tipo byte,
 dichiarati in quest'ordine.

- una funzione forward(turtle *Turtle, passi int)
che fa avanzare la tartaruga turtle nel verso della sua direzione
del numero di passi passato come parametro (non deve creare una nuova tartaruga)

- una funzione backward(turtle *Turtle, passi int)
che fa indietreggiare la tartaruga turtle rispetto al verso della sua direzione
del numero di passi passato come parametro (non deve creare una nuova tartaruga)

- una funzione setDir(turtle *Turtle, dir byte) error
che, se dir è una delle quattro direzione valide ('N', 'E', 'S' o 'O'),
imposta/modifica la direzione di marcia a dir per la tartaruga turtle (non deve creare una nuova tartaruga)
e inoltre restituisce nil,
altrimenti non modifica la direzione di marcia e restituisce un error (ad es. "non valid direction")

- un metodo String che produce una descrizione dello stato della tartaruga (posizione e direzione)
nel seguente formato:
(x,y) dir D
ad esempio:
(0,0) dir N

- una funzione main che
	- crea una Turtle in posizione (0,0) e diretta verso Nord
	- legge da standard input una serie di comandi (vedi sotto), uno per riga, e si ferma quando legge 'stop' su una nuova riga
	- per ogni comando:
		- invoca la funzione corrispondente, che modifica lo stato - posizione o direzione - della tartaruga
		- se il comando non è tra quelli riconosciuti (vedi sotto) stampa "invalid command"
		- stampa il nuovo stato della tartaruga

I tipi di possibili comandi sono tre (vedi anche esempio sotto):
F_n, dove n è un intero
B_n, dove n è un intero
S_d, dove d è una lettera tra 'N', 'E', 'S' e 'O' (le quattro direzioni Nord, Est, Sud, Ovest)

e corrispondono rispettivamente alle azioni
- avanza (forward) di n passi
- indietreggia (backward) di n passi
- imposta (setDir) a d la direzione di marcia

-----------------------------------

Esempio di esecuzione
---------------------

Data in input la sequenza di comandi:
F_5
S_E
F_12
S_O
B_3
F_15
S_S
B_5
stop

il programma stampa:
(0,5) dir N
(0,5) dir E
(12,5) dir E
(12,5) dir O
(15,5) dir O
(0,5) dir O
(0,5) dir S
(0,10) dir S

*/

package main

import "testing"

var prog = "./turtle"

func TestEsempio1(t *testing.T) {
	lanciaGenerica(t, prog, "F_5\nS_E\nF_12\nS_O\nB_3\nF_15\nS_S\nB_5\nstop\n", "(0,5) dir N\n(0,5) dir E\n(12,5) dir E\n(12,5) dir O\n(15,5) dir O\n(0,5) dir O\n(0,5) dir S\n(0,10) dir S\n")
}

func TestEsempio2(t *testing.T) {
	lanciaGenerica(t, prog, "B_3\nB_2\nS_E\nF_10\nD_E\nB_2\nstop\n", "(0,-3) dir N\n(0,-5) dir N\n(0,-5) dir E\n(10,-5) dir E\ninvalid command\n(10,-5) dir E\n(8,-5) dir E\n")
}

func TestBackward(t *testing.T) {
	turtle := Turtle{0, 0, 'N'}
	backward(&turtle, 2)
	if turtle.y != -2 || turtle.x != 0 {
		t.Fail()
	}
}

func TestForward(t *testing.T) {
	turtle := Turtle{0, 0, 'E'}
	forward(&turtle, 3)
	if turtle.x != 3 || turtle.y != 0 {
		t.Fail()
	}
}

func TestNeg(t *testing.T) {
	turtle := Turtle{0, 0, 'E'}

	forward(&turtle, -2)
	backward(&turtle, -2)

	if turtle.x != 0 || turtle.y != 0 {
		t.Fail()
	}
}

func TestSetDir(t *testing.T) {
	turtle := Turtle{0, 0, 'N'}
	setDir(&turtle, 'S')
	if turtle.y != 0 {
		t.Fail()
	}
	if turtle.dir != 'S' {
		t.Fail()
	}
}

func TestSetWrongDir(t *testing.T) {
	turtle := Turtle{0, 0, 'N'}
	err := setDir(&turtle, 'W')
	if err == nil {
		t.Fail()
	}
	if turtle.y != 0 {
		t.Fail()
	}
	if turtle.dir != 'N' {
		t.Fail()
	}
}
