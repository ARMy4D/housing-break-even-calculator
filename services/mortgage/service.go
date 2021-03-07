package mortgage

import "context"

const ErrorRange ErrorCode = 2000
const ServiceName string = "mortgage_service"

type Configuration struct {
}

type Service interface {
	CalculateYearlyMortgageIntrest(ctx context.Context, mortgageAmount float32, currentYear int, termInYears int, intrestRate float32) (totalPayment float32, err Errorer)
	CalculateYearlyMortgageRepayment(ctx context.Context, mortgageAmount float32, currentYear int, termInYears int) (totalPayment float32, err Errorer)
}
