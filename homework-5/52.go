package main

import (
//		"io/ioutil"
		"log"
		"os"	
		"fmt"
		"encoding/csv"
)



func main() {
	
	phonebook := [][]string{
		{"NAME", "PHONE"},
		{"Ivan", "89265550101"},
		{"Petr", "89165550202"},
		{"Michail", "89855550303"},
		{"Elena", "89175550404"},

	}

	fmt.Println(phonebook)

	file, err := os.Create("phonebook.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()


	w := csv.NewWriter(file) 
		defer w.Flush()


	for _, record := range phonebook {

	if err := w.Write(record); err != nil {
		log.Fatal(err)
	}
}


}