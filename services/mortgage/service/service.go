package mortgage

import (
	"context"
	"errors"

	"github.com/army4d/housing-break-even-calculator/services/mortgage"
	"github.com/go-kit/kit/log"
)

type mortService struct {
	logger log.Logger

	configuration mortgage.Configuration
}

func (s mortService) CalculateMortgagePaidAmount(mortgageAmount float32, currentYear int, termInYears int) (totalPayment float32, err mortgage.Errorer) {
	err = mortgage.ErrCalculation

	if mortgageAmount < 0 {
		err = err.AddToWrapped([]error{mortgage.ErrUnsupportedValue.Wrap([]error{errors.New("mortgage can't be lower than 0")})})
	}
	if currentYear < 1 {
		err = err.AddToWrapped([]error{mortgage.ErrUnsupportedValue.Wrap([]error{errors.New("current year can't be lower than 1")})})
	}
	if termInYears < 1 {
		err = err.AddToWrapped([]error{mortgage.ErrUnsupportedValue.Wrap([]error{errors.New("term can't be lower than 1")})})
	}

	if len(err.UnWrap()) > 0 {
		return 0, err
	}

	if currentYear > termInYears {
		return mortgageAmount, nil
	}

	return (mortgageAmount / float32(termInYears) * float32(currentYear-1)), nil
}

func (s mortService) CalculateYearlyMortgageIntrest(ctx context.Context, mortgageAmount float32, currentYear int, termInYears int, intrestRate float32) (totalPayment float32, err mortgage.Errorer) {
	err = mortgage.ErrCalculation

	if mortgageAmount < 0 {
		err = err.AddToWrapped([]error{mortgage.ErrUnsupportedValue.Wrap([]error{errors.New("mortgage can't be lower than 0")})})
	}
	if currentYear < 1 {
		err = err.AddToWrapped([]error{mortgage.ErrUnsupportedValue.Wrap([]error{errors.New("current year can't be lower than 1")})})
	}
	if termInYears < 1 {
		err = err.AddToWrapped([]error{mortgage.ErrUnsupportedValue.Wrap([]error{errors.New("term can't be lower than 1")})})
	}

	if len(err.UnWrap()) > 0 {
		return 0, err
	}
	if currentYear > termInYears {
		return 0, nil
	}
	if mortgagePaymentsDone, err := s.CalculateMortgagePaidAmount(mortgageAmount, currentYear, termInYears); err == nil {
		return ((mortgageAmount - mortgagePaymentsDone) * (intrestRate / 100)), nil
	} else {
		return 0, err
	}
}

func (s mortService) CalculateYearlyMortgageRepayment(ctx context.Context, mortgageAmount float32, currentYear int, termInYears int) (totalPayment float32, err mortgage.Errorer) {
	err = mortgage.ErrCalculation

	if mortgageAmount < 0 {
		err = err.AddToWrapped([]error{mortgage.ErrUnsupportedValue.Wrap([]error{errors.New("mortgage can't be lower than 0")})})
	}
	if termInYears < 1 {
		err = err.AddToWrapped([]error{mortgage.ErrUnsupportedValue.Wrap([]error{errors.New("term can't be lower than 1")})})
	}

	if len(err.UnWrap()) > 0 {
		return 0, err
	}
	if currentYear > termInYears {
		return 0, nil
	}
	return mortgageAmount / float32(termInYears), nil
}

func NewService(logger log.Logger) mortgage.Service {
	return &mortService{
		logger, mortgage.Configuration{},
	}
}
