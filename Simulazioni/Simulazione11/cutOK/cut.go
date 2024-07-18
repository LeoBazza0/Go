package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, j, file := leggiInput()
	var line string
	for h := 0; h < len(file); h++ {
		if len(file[h]) == 0 {
			fmt.Println()
			continue
		}
		for k := i - 1; k < j; k++ {
			if len(file[h]) < j && k == len(file[h]) && len(file[h]) != 0 {
				//	fmt.Println(line)
				//	line = ""
				break
			}
			line += file[h][k]
		}
		fmt.Println(line)
		line = ""
	}
}

// faccio una slice (una riga per cella) di slice (una lettera per cella)
func leggiInput() (i, j int, file [][]string) {
	filei, _ := os.Open(os.Args[3])
	scanner := bufio.NewScanner(filei)
	for scanner.Scan() {
		line := scanner.Text()
		files1 := strings.Split(line, "")
		file = append(file, files1)
	}
	i, _ = strconv.Atoi(os.Args[1])
	j, _ = strconv.Atoi(os.Args[2])
	filei.Close()
	return i, j, file
}
