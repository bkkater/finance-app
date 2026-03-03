package calc

import "math"

type InterestCalculator interface {
	CalculateCompound(capital float64, rate float64, months int) (float64, []float64)
}

type FinanceService struct{}

func (service *FinanceService) CalculateCompound(capital float64, rate float64, months int) (float64, []float64) {
	var monthlyInterests []float64
	currentBalance := capital

	for m := 1; m <= months; m++ {
		interestOfMonth := currentBalance * rate
		interestOfMonth = math.Round(interestOfMonth*100) / 100

		monthlyInterests = append(monthlyInterests, interestOfMonth)
		currentBalance += interestOfMonth
	}

	total := math.Round(currentBalance*100) / 100

	return total, monthlyInterests
}