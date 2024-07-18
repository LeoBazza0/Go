package main
import (
  "fmt"
  "strings"
  "os"
  "bufio"
  "strconv"
)

func main () {
  tot1:=1.0
  var tot float64
  scanner:=bufio.NewScanner (os.Stdin)
  for scanner.Scan() {
    line:=scanner.Text()
    sliceOrd:=strings.Split(line, "#")
    for i:=0; i<3; i++ {
      num, err := strconv.ParseFloat(sliceOrd[i], 64)
      if err!=nil {
        break
      }
      tot1*=num
        }
    tot+=tot1
  }
fmt.Println (tot)
}
