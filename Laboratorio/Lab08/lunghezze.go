package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lineString []string
var m map[int]int
var n int

func aggiornaConteggio(m map[int]int, line string)  {
  lineString = strings.Fields(line)
  for i:=0; i<len(lineString); i++ {
  	m[len(lineString[i])] += 1
  }
fmt.Print (m)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	m = map[int]int{}
	aggiornaConteggio (m, line)
}
