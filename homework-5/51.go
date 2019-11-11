package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)


// Проверим ошибки
func check(err error) {
	if err != nil { 
	log.Fatal(err)
}
}

func main() {
	file, err := os.Open("fileread.go")
	check(err)
	filename, err := file.Stat()
	check(err)
	fmt.Printf("File %v sucsessfuly opened\n", filename.Name())
	defer file.Close()

	stat, err := file.Stat()
	filesize := strconv.FormatInt(stat.Size(), 10)
	check(err)
	if filesize == "0" { 
		println("Empty file")
		os.Exit(1)
	} else {
		fmt.Printf("Size of file is %v bytes\n", filesize)
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	check(err)
	fmt.Printf("Sucsessfuly readed %v bytes of file\n", len(bs))


}
