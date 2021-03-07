package calculator

import (
	"context"
)

const ErrorRange ErrorCode = 1000
const ServiceName string = "calculator_server"

type Configuration struct {
	MinYearsToReside int
	MaxMortgageTerm  int
	MinMortgageTerm  int
}

type Service interface {
	Calculate(ctx context.Context, rent RentSetting, mortgage MortgageSetting, house HouseSetting) (breakEven BreakEven, err Errorer)
}

type ServiceMiddleware func(Service) Service
