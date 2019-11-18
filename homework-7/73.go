package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
		go readExit(conn)
		}


	}


	func readExit(c net.Conn) string {
		r := bufio.NewReader(c)
		msg, _ := r.ReadString('\n')
		if msg == "exit\r\n" {
			fmt.Println("trying to stop")
			os.Exit(0)
		}
	return msg
	}

func handleConn(c net.Conn) {

	defer c.Close()
	for {

				_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
				time.Sleep(2 * time.Second)
				if err != nil{
					log.Fatal(err)
				}

			}

		}