package tax

import "context"

const ErrorRange ErrorCode = 4000
const ServiceName string = "tax_service"

type Configuration struct {
}

type Service interface {
	CalculateYearlyTaxes(ctx context.Context, housePrice float32, currentYear int, taxes TaxSetting) (totalPayment float32, err Errorer)
}
