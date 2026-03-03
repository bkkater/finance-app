package rpctypes

type InvestmentRequest struct {
	InitialCapital float64
	MonthlyRate    float64
	Months         int
}

type InvestmentResponse struct {
	TotalAmount      float64
	MonthlyInterests []float64
}