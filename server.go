package main

import (
	"fmt"
	"main/handler"
	"main/types"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1234")

	handler.ErrorHandler(err)
	defer listener.Close()
	// defer itu utk ngejalanin code di atasnya

	for {
		conn, err := listener.Accept()
		// conn = connection utk nerima pesan dr client
		handler.ErrorHandler(err)
		go handleClient(conn)
		// ini go routine = supaya handle dr si client bisa jalan terus menerus
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	timeoutDuration := 5 * time.Second

	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		payload, err := types.Decode(conn)
		// payload itu buat ngedapetin hasil decode dr si client (krna si client ngirim pesan dlm bentuk bits)
		// decode itu pesan rahasia di balikin ke pesan aslinya

		handler.ErrorHandler(err)

		fmt.Println("Client : ", string(payload.Bytes()))

		var p types.Binary

		p = types.Binary("Recived: " + string(payload.Bytes()))

		_, err = p.WriteTo(conn)
		handler.ErrorHandler(err)

	}
}
