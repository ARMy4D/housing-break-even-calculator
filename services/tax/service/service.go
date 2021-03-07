package tax

import (
	"context"
	"errors"

	"github.com/army4d/housing-break-even-calculator/services/tax"
	"github.com/go-kit/kit/log"
)

type taxService struct {
	logger log.Logger

	configuration tax.Configuration
}

func (s taxService) CalculateYearlyTaxes(ctx context.Context, housePrice float32, currentYear int, taxes tax.TaxSetting) (totalPayment float32, err tax.Errorer) {
	err = tax.ErrCalculation

	if housePrice < 0 {
		err = err.AddToWrapped([]error{tax.ErrUnsupportedValue.Wrap([]error{errors.New("house price can't be lower than 0")})})
	}
	if currentYear < 1 {
		err = err.AddToWrapped([]error{tax.ErrUnsupportedValue.Wrap([]error{errors.New("current year can't be lower than 1")})})
	}

	if len(err.UnWrap()) > 0 {
		return 0, err
	}

	if currentYear == 1 {
		totalPayment += housePrice * (taxes.PropertyTransferTaxRate / 100)
	}

	totalPayment += housePrice * (taxes.PropertyTaxRate / 100)

	return totalPayment, nil
}

func NewService(logger log.Logger) tax.Service {
	return &taxService{
		logger, tax.Configuration{},
	}
}
