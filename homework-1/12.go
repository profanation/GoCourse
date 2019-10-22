package main

import (
	"fmt"
	"math"
)
func main() {
	var a float64
	var b float64
	fmt.Print("Введите величину катета а:")
	fmt.Scan(&a)
	fmt.Print("Введите величину катета b:")
	fmt.Scan(&b)
	s := (a*b)/2
	c := math.Sqrt(a*a + b*b)
	h := (a*b)/c
	fmt.Printf("s= %.2f\nc= %.2f\nh= %.2f\n", s, c, h)

}