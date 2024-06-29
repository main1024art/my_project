package main

import (
	"go_code/rpc/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(&service.Caluclator{})
	conn, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer conn.Close()
	log.Println("listening on :1234")
	for {
		conn, err := conn.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
