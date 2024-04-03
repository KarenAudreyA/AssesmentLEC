package main

import (
	"io"
	"main/handler"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9999")

	handler.ErrorHandler(err)

	for {
		conn, err := listener.Accept()

		handler.ErrorHandler(err)

		go handleServer(conn)
	}
}

func proxyForward(from io.Reader, to io.Writer) error {

	fromWriter, fromIsWritter := from.(io.Writer)
	toReader, toIsReader := to.(io.Reader)

	if fromIsWritter && toIsReader {
		go func() {
			_, err := io.Copy(fromWriter, toReader)
			handler.ErrorHandler(err)
			return
		}()
	}

	_, err := io.Copy(to, from)
	handler.ErrorHandler(err)

	return err
}

func handleServer(from net.Conn) {

	defer from.Close()

	to, err := net.Dial("tcp", "localhost:1234")
	handler.ErrorHandler(err)

	// from -> client
	// to -> server
	// misalkan server kirim balik ke client
	// from -> server
	// to -> client
	// dihandle oleh function goroutine
	err = proxyForward(from, to)

	handler.ErrorHandler(err)
}
