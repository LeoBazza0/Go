package main

import "fmt"

const STOP = 9999

func main() {
	var dayRain, totRain, count, countGiorni int
	var averageRainfall float64
	var slicePioggia []int
	fmt.Print("mm di pioggia: ")
	for {
		fmt.Scan(&dayRain)
		if dayRain == STOP {
			break
		}
		slicePioggia = append(slicePioggia, dayRain)
		count++
		totRain += dayRain
	}
	if count == 0 {
		fmt.Println("nessun dato disponibile")
	} else {
		averageRainfall = float64(totRain) / float64(count)
		fmt.Println("media:", averageRainfall)
	}
	for i := 0; i < len(slicePioggia); i++ {
		if float64(slicePioggia[i]) > averageRainfall {
			countGiorni++
		}
	}
	fmt.Println(countGiorni)

	for i := len(slicePioggia) - 1; i >= 0; i-- {
		if slicePioggia[i] != 0 {
			fmt.Println("l'ultimo giorno che ha piovuto è stato il giorno numero: ", i+1)
			break
		}
	}
}
