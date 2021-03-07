package rent

import "context"

const ErrorRange ErrorCode = 3000
const ServiceName string = "rent_service"

type Configuration struct {
	MonthsInAYear int
}

type Service interface {
	CalculateYearlyRent(ctx context.Context, rentAmount float32, currentYear int, increaseRate float32) (totalPayment float32, err Errorer)
}
