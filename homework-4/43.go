
package main


import "fmt"


func Help() {
	fmt.Println("USAGE:")
	fmt.Println("\t[arg1] [+-*/func] [[arg2]]")
	fmt.Printf("\nfunc:\n \t- sqrt, args:1\n \t- log, args:2\n")
	fmt.Printf("\t- abs, args:1\n \t- log, args:1\n")
	fmt.Printf("\t- ln, args:1\n \t- sin, args:1\n")
	fmt.Printf("\t- tan, args:1\n \t- arcsin, args:1\n")
	fmt.Printf("\t- arccos, args:1\n \t- arctan, args:1\n")
	fmt.Printf("\t- max, args:2\n \t- min, args:2\n")
}





func main() {
	Help()
	
	// input := ""
    // for {
    //     fmt.Print("> ")
    //     if _, err := fmt.Scanln(&input); err != nil {
    //         fmt.Println(err)
    //         continue
    //     }

    //     if input == "exit" {
    //         break
    //     }

	// 	if input == "help" {
	// 		Help()
	// 	}

    //     if res, err := calculator.Calculate(input); err == nil {
    //         fmt.Printf("Результат: %v\n", res)
    //     } else {
    //         fmt.Println("Не удалось произвести вычисление")
    //     }
    // }
}
