package main

import (
	"fmt"
	"os"
//	"time"
)

func main() {
	cancelCh:=make(chan struct{})

	go scanCancl(cancelCh)

	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x:=0; ;x++ {
//
			naturals <- x

		}

	}()

	go func() {
			for x:= range naturals{
				select {

				case <-naturals:
					squares <- x * x

				case <-cancelCh:
					close(naturals)

				}

			}
			close(squares)
			}()

	// печать
//go func() {
	for range squares{
		fmt.Println(<-squares)
	}
//fmt.Scanln()

//}()

}

func scanCancl(cancel chan<- struct{}) {
	os.Stdin.Read(make([]byte, 1))
	cancel <- struct{}{}
}