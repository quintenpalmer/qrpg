package qrpg

import (
	"net"
	"fmt"
)

const (
	RECV_BUF_LEN = 1024
)

func ListenForever(port string, model *Model) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("Could not listen on port :" + port)
	}
	fmt.Println("Listening on port :" + port)
	running := true
	for running {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		running = handleConnection(conn, model)
	}
}

func handleConnection(conn net.Conn, model *Model) bool {
	fmt.Println("Handling a connection")
	buf := make([]byte, RECV_BUF_LEN)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Write([]byte("error\n"))
		fmt.Println(err)
		return true
	}
	buf = buf[0:n]
	response, err := Respond(string(buf), model)
	if err != nil {
		conn.Write([]byte("505\n"))
	}
	if response == "exit" {
		return false
	} else {
		conn.Write([]byte(response + "\n"))
		return true
	}
}
