package main 

import (
	"net"
	"fmt"
	"os"
)

func main() {
	service := ":1300"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	//fmt.Println(tcpAddr)

	server, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close()
	}
}

func handleClient(conn net.Conn) {
	var buf[512]byte 
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return 
		}
		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err!= nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}