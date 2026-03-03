package main

import (
	"fmt"
	"log"
	"net/rpc"

	rpctypes "github.com/bkkater/finance-rpc-service/pkg/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:3333")

	if err != nil {
		log.Fatal("Erro ao conectar no servidor:", err)
	}

	var req rpctypes.InvestmentRequest

	fmt.Println("--- Sistema de Investimentos F$ ---")
	fmt.Print("Valor do investimento inicial (R$): ")
	fmt.Scanln(&req.InitialCapital)

	fmt.Print("Taxa de juros mensal (ex: 0.01 para 1%): ")
	fmt.Scanln(&req.MonthlyRate)

	fmt.Print("Período do investimento (meses): ")
	fmt.Scanln(&req.Months)

	var res rpctypes.InvestmentResponse
	err = client.Call("FinanceHandler.Calculate", req, &res)

	if err != nil {
		log.Fatal("Erro na chamada:", err)
	}

	fmt.Printf("\nResultado: R$ %.2f\n", res.TotalAmount)
	fmt.Println("Rendimento Mensal:")
	for i, v := range res.MonthlyInterests {
		fmt.Printf("   Mês %d: R$ %.2f\n", i+1, v)
	}
}