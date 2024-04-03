package main

// client itu pasti kirim ke proxy bukan lngsng ke server

import (
	"bufio"
	"fmt"
	"main/handler"
	"main/types"
	"net"
	"os"
	"time"
)

var choice int

func main() {
	dial, err := net.Dial("tcp", "localhost:9999")

	handler.ErrorHandler(err)

	defer dial.Close()
	scanner := bufio.NewScanner(os.Stdin)
	var message string

	for {
		printMenu()
		switch choice {
		case 1:
			fmt.Println("Send Message: ")
			message = scanner.Text()

			data := types.Binary(message)
			_, err := data.WriteTo(dial)
			handler.ErrorHandler(err)

			// menerima data yang disend oleh server (proxy)
			// set deadline selama 5 detik
			// melewati 5 detik maka koneksi putus

			dial.SetReadDeadline(time.Now().Add(5 * time.Second))

			// cara ngeread
			p, err := types.Decode(dial)

			// menghandle deadline
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					fmt.Println("Time out")
					return
				} else {
					handler.ErrorHandler(err)
				}
			}

			fmt.Println("Server: " + string(p.Bytes()))

		case 2:
			return
		}
	}
}

func printMenu() {
	fmt.Println("Hii Welcome..")
	fmt.Println("1. Send a message")
	fmt.Println("2. Exit")
	fmt.Print(">> ")
	fmt.Scanf("%d\n", &choice)
}
