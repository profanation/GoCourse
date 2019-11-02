package main

import (
	"fmt"
	"os"
	"strconv"
)

func isEven(values int64){
		if values % 2 == 0 {
			fmt.Println("Вы ввели четное число")
		} else {
			fmt.Println("Вы ввели нечетное число")
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
	isEven(num)

}