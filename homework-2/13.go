package main

import "fmt"

func main() {
	var fn1,fn2 uint64
	fn1 = 1
	fn2 =1
	fmt.Print("0\n1\n1\n")
	for i:=3; i<=100; i++  {
		fntmp := fn1+fn2
		fn1 = fn2
		fn2 = fntmp
		fmt.Println(fntmp)

	}

}