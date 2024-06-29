package main

import (
	"fmt"
	"go_code/rpc_client/data"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	request := &data.Caluclatorrequest{2, 5}
	response := &data.Caluclatorresponse{}
	err = conn.Call("Caluclator.Add", request, response)
	fmt.Println(response.Result, "11111")
}
