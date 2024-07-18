package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var altobasso string
	var countTrue, countFalse int
	rand.Seed(time.Now().UnixNano())
	primonumero := rand.Int()
	for {
		fmt.Print("alto/basso? ")
		fmt.Scan(&altobasso)
		if altobasso == "0" {
			break
		}
		numrand := rand.Int()
		if primonumero == numrand {
			numrand++
		}
		if altobasso == "alto" && primonumero < numrand || altobasso == "basso" && primonumero < numrand {
			countTrue++
			fmt.Println("giusto")
		} else {
			countFalse++
			fmt.Println("sbagliato")
		}
		fmt.Println(primonumero, numrand, countTrue, countFalse)
		primonumero = numrand
	}
	if countTrue > countFalse {
		fmt.Println("hai vinto")
	} else if countTrue < countFalse {
		fmt.Println("hai perso")
	} else {
		fmt.Println("pari")
	}
	fmt.Printf("Vittorie: %.2f %%\n", float64(countTrue)/float64(countTrue+countFalse)*100)
	fmt.Printf("Sconfitte: %.2f %%\n", float64(countFalse)/float64(countTrue+countFalse)*100)
}
