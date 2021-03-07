package calculator

import (
	"context"
	"errors"

	"github.com/army4d/housing-break-even-calculator/services/calculator"
	"github.com/go-kit/kit/log"
)

type validationMiddleware struct {
	logger log.Logger

	configuration calculator.Configuration

	next calculator.Service
}

func (mw validationMiddleware) Calculate(ctx context.Context, rent calculator.RentSetting, mortgage calculator.MortgageSetting, house calculator.HouseSetting) (breakEven calculator.BreakEven, err calculator.Errorer) {
	err = calculator.ErrConstraintReached
	if mw.configuration.MinYearsToReside != 0 && house.YearsExpectedToReside < mw.configuration.MinYearsToReside {
		err = err.AddToWrapped([]error{calculator.ErrMinReside})
	}
	if mw.configuration.MaxMortgageTerm != 0 && mortgage.Term > mw.configuration.MaxMortgageTerm {
		err = err.AddToWrapped([]error{calculator.ErrMaxMortgageTerm})
	}

	if mw.configuration.MinMortgageTerm != 0 && mortgage.Term < mw.configuration.MinMortgageTerm {
		err = err.AddToWrapped([]error{calculator.ErrMinMortgageTerm})
	}

	if mortgage.DownPayment < 0 {
		err = err.AddToWrapped([]error{errors.New("down payment can't be lower than 0")})
	}

	if len(err.UnWrap()) > 0 {
		return
	}

	return mw.next.Calculate(ctx, rent, mortgage, house)
}

func NewValidationMiddleware(logger log.Logger, minYearsToReside, minMortgageTerm, maxMortgageTerm int) calculator.ServiceMiddleware {
	return func(next calculator.Service) calculator.Service {
		return validationMiddleware{logger,
			calculator.Configuration{
				MinYearsToReside: minYearsToReside,
				MinMortgageTerm:  minMortgageTerm,
				MaxMortgageTerm:  maxMortgageTerm,
			},
			next}
	}
}
