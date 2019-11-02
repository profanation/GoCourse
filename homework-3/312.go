package main

import (
		"fmt"

		)
		

func main() {

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

	var lada = Sedan{"lada", 1989, 100, false, true, 0}

	var gazel = Miniven{"gazel",1995,300,false,true,0}

	fmt.Println(lada)
	fmt.Println(gazel)


	// Разгрузим машины
	lada.FreeVolume=100
	gazel.FreeVolume=300

	fmt.Println(lada)
	fmt.Println(gazel)

}