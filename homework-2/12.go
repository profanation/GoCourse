package main

import (
	"fmt"
	"strconv"
	"os"
)

func is3ism(values int64){
	if values % 3 == 0 {
		fmt.Println("Введенное число делится на 3")
	} else {
		fmt.Println("Введенное число на 3 не делится")
	}

}

func main()  {
	a := "Введите число: "
	//fmt.Println("Введте число")
	fmt.Print(a)
	var input string
	fmt.Scan(&input)
	num, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		fmt.Println("Вы ввели что-то не то...")
		fmt.Println(err)
		os.Exit(1)
	}
	is3ism(num)

}