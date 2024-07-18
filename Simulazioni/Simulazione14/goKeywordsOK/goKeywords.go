package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	mappa := make(map[string]int)
	sliceKey := []string{"break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for", "func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select", "struct", "switch", "type", "var"}
	if len(os.Args) < 2 {
		fmt.Println("A file name, please")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File not found")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sliceLine := strings.Split(line, " ")
		for i := 0; i < len(sliceLine); i++ {
			for j := 0; j < len(sliceKey); j++ {
				if sliceLine[i] == sliceKey[j] {
					mappa[sliceKey[j]]++
				}
			}
		}
	}
	printKeywords(mappa)
}

func printKeywords(mappa map[string]int) {
	keys := make([]string, 0, len(mappa))
	for k, _ := range mappa {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if mappa[k] != 0 {
			fmt.Println(k, mappa[k])
		} else {
			continue
		}
	}

}

/*
"break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for",
"func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select",
"struct", "switch", "type", "var"
*/
