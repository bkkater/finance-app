package main

import (
	"fmt"
	"time"

	rpctypes "github.com/bkkater/finance-rpc-service/pkg/rpc"
	"github.com/bkkater/finance-rpc-service/services/calc"
)

type FinanceHandler struct {
	Calculator calc.InterestCalculator
}

func (handler *FinanceHandler) Calculate(args rpctypes.InvestmentRequest, reply *rpctypes.InvestmentResponse) error {
	start := time.Now()

	total, monthly := handler.Calculator.CalculateCompound(args.InitialCapital, args.MonthlyRate, args.Months)

	reply.TotalAmount = total
	reply.MonthlyInterests = monthly

	fmt.Printf("[LOG] Processamento concluído em: %v\n", time.Since(start))
	return nil
}
