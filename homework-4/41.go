package main

import "fmt"

type Sedan struct {
	Name       string
	Year       int
	Volume     int
	IsEgnition bool
	IsWindOpen bool
	FreeVolume int
}

type Miniven struct {
	Name       string
	Year       int
	Volume     int
	IsEgnition bool
	IsWindOpen bool
	FreeVolume int
}


// func (f *Sedan) unloadvol(FreeVolume int) {
// 	f.FreeVolume=0

// }

func (m Sedan) Loadedvol() int {
	return m.Volume-m.FreeVolume
} 


func (m Miniven) Loadedvol() int {
	return m.Volume-m.FreeVolume
}


type CurrLoad interface {
	Loadedvol() int
}

func  SummLoad(cars ...CurrLoad) int {

	ld:=0
	for _, cars := range(cars) {
		//ld += CurrLoad.loadedvol(loads)
		ld += cars.Loadedvol()
	}
	return ld
} 


func main() {

 var lada = Sedan{"lada", 1989, 100, false, true, 0}

 var gazel = Miniven{"gazel",1995,300,false,true,0}




A:=SummLoad(lada, gazel)


fmt.Println(lada)

fmt.Println(gazel)

fmt.Println(A)

	
}