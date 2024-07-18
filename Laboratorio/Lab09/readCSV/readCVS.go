

package main
import (
  "fmt"
  "io"
)

type rilevazione struct {
  year, month, day, hour, min  int
  temperatura, umidita float64
}
func main () {
  var rilev, minTemp, maxTemp, minHumid, maxHumid rilevazione
    _, err:=fmt.Scanf("%d-%d-%d/%d:%d,%f,%f", &rilev.year, &rilev.month, &rilev.day, &rilev.hour, &rilev.min, &rilev.temperatura, &rilev.umidita )
    if err != io.EOF {
		minTemp = rilev
		maxTemp = rilev
		minHumid = rilev
		maxHumid = rilev
	}
for {
  _, err:=fmt.Scanf("%d-%d-%d/%d:%d,%f,%f", &rilev.year, &rilev.month, &rilev.day, &rilev.hour, &rilev.min, &rilev.temperatura, &rilev.umidita )
  if err==io.EOF {
    break
  }
  if minTemp.temperatura>rilev.temperatura {
    minTemp=rilev
  }
  if maxTemp.temperatura<rilev.temperatura {
    maxTemp=rilev
  }
  if minHumid.umidita>rilev.umidita {
    minHumid=rilev
  }
  if maxHumid.umidita<rilev.umidita {
    maxHumid=rilev
  }
}
fmt.Printf("minTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", minTemp.temperatura, minTemp.hour, minTemp.min, minTemp.day, minTemp.month, minTemp.year)

	fmt.Printf("maxTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", maxTemp.temperatura, maxTemp.hour, maxTemp.min, maxTemp.day, maxTemp.month, maxTemp.year)

	fmt.Printf("minHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", minHumid.umidita, minHumid.hour, minHumid.min, minHumid.day, minHumid.month, minHumid.year)

	fmt.Printf("maxHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", maxHumid.umidita, maxHumid.hour, maxHumid.min, maxHumid.day, maxHumid.month, maxHumid.year)
}


/*
package main

import (
	"fmt"
	"io"
)

type Rilevazione struct {
	yy, mm, dd, hh, min int
	temp, humid         float64
}

func main() {
	var rilev, minTemp, maxTemp, minHumid, maxHumid Rilevazione

	_, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f", &rilev.yy, &rilev.mm, &rilev.dd, &rilev.hh, &rilev.min, &rilev.temp, &rilev.humid)
	//fmt.Println(minTemp, maxTemp, minHumid, maxHumid)

	if err != io.EOF {
		minTemp = rilev
		maxTemp = rilev
		minHumid = rilev
		maxHumid = rilev
	}

	for {
		_, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f", &rilev.yy, &rilev.mm, &rilev.dd, &rilev.hh, &rilev.min, &rilev.temp, &rilev.humid)
		if err == io.EOF {
			break
		}
		if rilev.temp > maxTemp.temp {
			maxTemp = rilev
		} else if rilev.temp < minTemp.temp {
			minTemp = rilev
		}
		if rilev.humid > maxHumid.humid {
			maxHumid = rilev

		} else if rilev.humid < minHumid.humid {
			minHumid = rilev
		}
	}

	fmt.Printf("minTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", minTemp.temp, minTemp.hh, minTemp.min, minTemp.dd, minTemp.mm, minTemp.yy)

	fmt.Printf("maxTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", maxTemp.temp, maxTemp.hh, maxTemp.min, maxTemp.dd, maxTemp.mm, maxTemp.yy)

	fmt.Printf("minHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", minHumid.humid, minHumid.hh, minHumid.min, minHumid.dd, minHumid.mm, minHumid.yy)

	fmt.Printf("maxHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", maxHumid.humid, maxHumid.hh, maxHumid.min, maxHumid.dd, maxHumid.mm, maxHumid.yy)

}
/*
package main

import (
	"fmt"
	"io"
)

type TimeStamp struct {
	anno, mese, giorno, ora, minuto int
}

func (ts *TimeStamp) String() string {
	return fmt.Sprintf("misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d", ts.ora, ts.minuto, ts.giorno, ts.mese, ts.anno)
}

type Misurazione struct {
	timestamp TimeStamp
	temperatura float64
	umidita float64
}
//Es: 1.0° misura delle ore 4, minuti 22 del giorno 11 del mese 12 dell'anno 2022
func (m *Misurazione) StringTemperatura() string {
	return fmt.Sprintf("%.1f° %s", m.temperatura, m.timestamp.String())
}

func (m *Misurazione) StringUmidita() string {
	return fmt.Sprintf("%.1f%% %s", m.umidita, m.timestamp.String())
}

func main() {
	minTemp := Misurazione{TimeStamp{}, 1000, 0}
	maxTemp := Misurazione{TimeStamp{}, -1000, 0}
	minHumid := Misurazione{TimeStamp{}, 0, 1000}
	maxHumid := Misurazione{TimeStamp{}, 0, -1000}
	var anno, mese, giorno, ora, minuto int
	var temperatura, umidita float64
	for {
		_, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f", &anno, &mese, &giorno, &ora, &minuto, &temperatura, &umidita)
		if err == io.EOF {
			break
		}
		if temperatura > maxTemp.temperatura {
			maxTemp = Misurazione{timestamp: TimeStamp{anno: anno, mese: mese, giorno: giorno, ora: ora, minuto: minuto}, temperatura: temperatura, umidita: umidita}
		} else if temperatura < minTemp.temperatura {
			minTemp = Misurazione{timestamp: TimeStamp{anno: anno, mese: mese, giorno: giorno, ora: ora, minuto: minuto}, temperatura: temperatura, umidita: umidita}
		}
		if umidita > maxHumid.umidita {
			maxHumid = Misurazione{timestamp: TimeStamp{anno: anno, mese: mese, giorno: giorno, ora: ora, minuto: minuto}, temperatura: temperatura, umidita: umidita}
		} else if umidita < minHumid.umidita {
			minHumid = Misurazione{timestamp: TimeStamp{anno: anno, mese: mese, giorno: giorno, ora: ora, minuto: minuto}, temperatura: temperatura, umidita: umidita}
		}
	}
	fmt.Println("minTemp:", minTemp.StringTemperatura())
	fmt.Println("maxTemp:", maxTemp.StringTemperatura())
	fmt.Println("minHumid:", minHumid.StringUmidita())
	fmt.Println("maxHumid:", maxHumid.StringUmidita())

}

*/
