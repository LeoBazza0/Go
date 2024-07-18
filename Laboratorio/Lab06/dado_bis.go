package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const nFacce = 6
	var lancio, n int
	var frequenza [nFacce + 1]int

	fmt.Scan(&n)

	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		lancio = rand.Intn(nFacce) + 1
		frequenza[lancio] += 1
	}

	for i := 0; i < 6; i++ {
		fmt.Println(i+1, ":", frequenza[i+1], "(", frequenza[i+1]*100/n, "%)")
	}
}
