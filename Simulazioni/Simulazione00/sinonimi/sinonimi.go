package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("2 file names, please")
		return
	}
	sinonimi, err1 := os.Open(os.Args[1])
	file2, err2 := os.Open(os.Args[2])
	if err1 != nil {
		fmt.Println("file 1 not found")
		return
	} else if err2 != nil {
		fmt.Println("file 2 not found")
		return
	}
	scanner := bufio.NewScanner(file2)
	stringSinonimi := takeFirst(sinonimi)
	for scanner.Scan() {
		line := scanner.Text()
		sliceString := strings.Split(line, ":")
		if strings.Contains(sliceString[i], stringSinonimi)
	}
}

func takeFirst(sinonimi *os.File) (stringaPrimi string) {
	scanner := bufio.NewScanner(sinonimi)
	for scanner.Scan() {
		line := scanner.Text()
		sliceLine := strings.Split(line, ":")
		stringaPrimi += sliceLine[0]
	}
	return
}

func parolaConSinonimi(parola string, sinonimi *os.File) string {

}
