package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/bkkater/finance-rpc-service/services/calc"
)

func main() {
	service := &calc.FinanceService{}
	handler := &FinanceHandler{
		Calculator: service,
	}

	rpc.Register(handler)
	listener, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal("Erro no servidor:", err)
	}

	defer listener.Close()
	fmt.Println("Servidor Online [Porta 3333]")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("Erro na conexão:", err)
			continue
		}
		fmt.Printf("Conexão recebida de: %s\n", connection.RemoteAddr().String())
		go rpc.ServeConn(connection)
	}
}