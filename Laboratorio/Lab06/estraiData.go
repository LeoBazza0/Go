package main

import (
   "fmt"
   "strings"
   "strconv"
)

var dataGMA string


func main() {
   fmt.Print("data nel formato giorno/mese/anno: ")
   fmt.Scan(&dataGMA)
   giorniInt, mesiInt, anniInt := stringGMA2GMA(dataGMA)
   fmt.Println("giorno", giorniInt)
   fmt.Println("mese", num2mese(mesiInt))
   fmt.Println("anno", anniInt)

}

func stringGMA2GMA(data string) (int, int, int) {
  var giorniS, mesiS, anniS string

  slash1:=strings.Index(dataGMA, "/")
  for i:=0; i<slash1; i++ {
giorniS+=string(dataGMA[i])
  }
  giorniInt, err1 := strconv.Atoi(giorniS)
  if err1 != nil {
		fmt.Print("problema")
	}

slash2:=strings.LastIndex(data, "/")
  for j:=(slash1+1); j<slash2; j++ {
    mesiS+=string(dataGMA[j])
  }
  mesiInt, err2 := strconv.Atoi(mesiS)
  if err2 != nil {
		fmt.Print("problema")
	}

for k:=(slash2+1); k<len(dataGMA); k++ {
  anniS+=string(dataGMA[k])
}
anniInt, err3 :=strconv.Atoi (anniS)
if err3 != nil {
  fmt.Print("problema")
}
  return giorniInt, mesiInt, anniInt
 }

func num2mese(mesiInt int) string {
   var mesi = [13]string{"", "gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"}
   return mesi[mesiInt]
}
