package main


import (
		"fmt"	
		)

func main()  {
	const usd float32 = 63.72	
	fmt.Print("Введите колличество рублей:")
	var rub float32
	fmt.Scan(&rub)
	rub2usd := rub/usd
	fmt.Printf("%v рублей равны %.2f долларов США\n", rub , rub2usd)
}
