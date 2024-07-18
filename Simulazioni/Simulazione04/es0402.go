package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Comando struct {
	direzione string
	passi     int
}

func main() {
	comandi := LeggiComandi()
	comanditot := AnalizzaComandi(comandi)
	stampaRisultati(comandi, comanditot)
}

func LeggiComandi() (comandi []Comando) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		sliceC := strings.Split(line, " ")
		num, _ := strconv.Atoi(sliceC[1])
		comandi = append(comandi, Comando{sliceC[0], num})
	}
	return comandi

}

func AnalizzaComandi(comandi []Comando) map[string]int {
	comanditot := make(map[string]int)
	for _, comando := range comandi {
		comanditot[comando.direzione] += comando.passi
	}
	return comanditot
}

func stampaRisultati(comandi []Comando, comanditot map[string]int) {
	fmt.Println("Movimenti totali:")
	for key, cont := range comanditot {
		fmt.Printf("%s %d\n", key, cont)
	}
	fmt.Println("Direzione in cui il robot deve compiere il minor numero di passi:")
	fmt.Println(min(comanditot, comandi))
	fmt.Println("Comandi Inversi:")
	for i := len(comandi) - 1; i > 0; i-- {
		fmt.Printf("%v, ", comandi[i].Inverso().String())
	}
	fmt.Printf("%v\n", comandi[0].Inverso().String())
}

func min(comanditot map[string]int, comandi []Comando) string {
	var min int = comanditot[comandi[0].direzione]
	var minKey string = comandi[0].direzione
	for key, cont := range comanditot {
		if cont < min {
			min = cont
			minKey = key
		}
	}
	return minKey
}

func (c Comando) String() string {
	return fmt.Sprintf("%s %d", c.direzione, c.passi)
}

func (c Comando) Inverso() (inverso Comando) {
	switch c.direzione {
	case "NORD":
		inverso.direzione = "SUD"
	case "SUD":
		inverso.direzione = "NORD"
	case "EST":
		inverso.direzione = "OVEST"
	case "OVEST":
		inverso.direzione = "EST"
	}
	inverso.passi = c.passi
	return
}

/*
OVEST 2
NORD 1
EST 5
SUD 4
EST 5
NORD 3
SUD 1
OVEST 3
*/
