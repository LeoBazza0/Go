/*
package main
import (
  "fmt"
  "time"

)

func main () {
  var hh, mm, ss int
  var myTime Clock
  fmt.Print("Inserisci un orario (formato hh:mm:ss) ")
  fmt.Scanf ("%d:%d:%d", &hh, &mm, &ss)
  fmt.Printf ("time (%d:%d:%d) \n", hh, mm, ss)
  myTime= Clock{hh, mm, ss}
  for myTime.hour >=0 && myTime.min>=0 && myTime.sec>0 {
    countdown(&myTime)
    fmt.Println(myTime)
    time.Sleep(time.Duration(1) * time.Second)
  }
}

type Clock struct {
  hour, min, sec int
}

func countdown (c *Clock) {
  c.sec--
  if c.sec<0 {
    updateMin(c)
    c.sec=59
  }
}

func updateMin (c *Clock) {
  c.min --
  if c.min<0 {
    updateHour(c)
    c.min=59
  }
}

func updateHour (c *Clock) {
  c.hour--
}
*/
package main

import (
	"fmt"
	"time"
)

type Clock struct {
	hour, min, sec int
}

func main() {
	var myCountdown Clock
	var h, m, s int
	fmt.Print("orario di partenza? (h m s) ")
	fmt.Scan(&h, &m, &s)
	myCountdown = Clock{h, m, s}
	for myCountdown.hour >= 0 && myCountdown.min >= 0 && myCountdown.sec >= 0 {
		fmt.Println(myCountdown)
		countdown(&myCountdown)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func countdown(c *Clock) {
	c.sec--
	if c.sec < 0 {
		updateMin(c)
		c.sec = 59
	}
}

func updateMin(c *Clock) {
	c.min--
	if c.min < 0 {
		updateHour(c)
		c.min = 59
	}
}

func updateHour(c *Clock) {
	c.hour--
}
