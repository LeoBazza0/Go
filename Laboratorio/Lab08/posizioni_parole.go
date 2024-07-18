package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)
func main () {
  var line string
  scanner:=bufio.NewScanner (os.Stdin)
  for scanner.Scan () {
    line+=scanner.Text ()
  }
  sliceStr:=strings.Fields(line)
  var mappa map[string]int
//  mappa = map[string]int
  for i:=0; i<len(sliceStr); i++ {
    mappa[sliceStr[i]]++
  }
  fmt.Println (mappa)
}
