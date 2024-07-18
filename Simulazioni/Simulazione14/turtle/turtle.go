package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Turtle struct {
	x, y int  //coordinate nel pinao cartesiano
	dir  byte // N, E, S O
}

func main() {
	tartaruga := Turtle{0, 0, 'N'}
	for {
		var stringa string
		fmt.Scan(&stringa)
		if stringa == "stop" {
			os.Exit(0)
			break
		} else {
			stringaSlice := strings.Split(stringa, "_")
			FBS := stringaSlice[0] //saranno stringhe
			NND := stringaSlice[1] //saranno int int byte
			if !strings.ContainsAny(FBS, "FBS") {
				fmt.Println("invalid command")
				continue
			}
			if FBS == "F" {
				passi, _ := strconv.Atoi(NND)
				forward(&tartaruga, passi)
			} else if FBS == "B" {
				passi, _ := strconv.Atoi(NND)
				backward(&tartaruga, passi)
			} else {
				byteDir := []byte(NND)
				setDir(&tartaruga, byteDir[0])
			}
			fmt.Println(tartaruga)
		}
	}
}

// molto probabilmente sbagliato
func forward(turtle *Turtle, passi int) {
	NESO := []byte("NESO")
	if &turtle.dir == &NESO[0] {
		turtle.y += passi
	} else if &turtle.dir == &NESO[2] {
		turtle.y -= passi
	} else if &turtle.dir == &NESO[1] {
		turtle.x += passi
	} else if &turtle.dir == &NESO[3] {
		turtle.x -= passi
	}
}

// molto probabilmente sbagliato
func backward(turtle *Turtle, passi int) {
	NESO := []byte("NESO")
	if &turtle.dir == &NESO[0] {
		turtle.y -= passi
	} else if &turtle.dir == &NESO[2] {
		turtle.y += passi
	} else if &turtle.dir == &NESO[1] {
		turtle.x -= passi
	} else if &turtle.dir == &NESO[3] {
		turtle.x += passi
	}
}

func setDir(turtle *Turtle, dir byte) error {
	//no comment aaaa
	return nil
}
