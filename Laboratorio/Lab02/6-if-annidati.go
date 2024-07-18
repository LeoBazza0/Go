package main
import "fmt"
func main() {
	/*
 dichiarare 4 costanti, secondo le quali si differenziano le variabili: si distinguono il RETTANGOLO e il TRIANGOLO e il PERIMETRO e l'AREA
 si scannerizza prima la figura (RETTANGOLO o TRIANGOLO) e poi l'operazione (PERIMETRO o AREA)
 se l'utente sceglie il RETTANGOLO si verifica se sceglie il PERIMETRO e si scrive la formula del perimetro, altrimenti si scrive la formula dell'AREA
 se l'utente invece non sceglie il RETTANGOLO (e quindi sceglie il TRIANGOLO), si verifica se sceglie il PERIMETRO e si scrive la formula del perimetro, altrimenti si scrive la formula dell'AREA
	*/
	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	)
	var figura, operazione int

	fmt.Print("rettangolo (1) o triangolo (2)? ")
	fmt.Scan(&figura)
	fmt.Print("perimetro (1) o area (2)? ")
	fmt.Scan(&operazione)

	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2")
		} else { //operazione == AREA
			fmt.Println("formula = base * altezza")
		}
	} else { //figura == TRIANGOLO
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ")
		} else { //operazione == AREA
			fmt.Println("formula = (base * altezza)/2")
		}
	}
}
