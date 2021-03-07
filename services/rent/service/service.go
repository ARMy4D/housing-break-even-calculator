package rent

import (
	"context"
	"errors"

	"github.com/army4d/housing-break-even-calculator/services/rent"
	"github.com/go-kit/kit/log"
)

type rentService struct {
	logger log.Logger

	configuration rent.Configuration
}

func (s rentService) CalculateYearlyRent(ctx context.Context, rentAmount float32, currentYear int, increaseRate float32) (totalPayment float32, err rent.Errorer) {
	err = rent.ErrCalculation

	if rentAmount < 0 {
		err = err.AddToWrapped([]error{rent.ErrUnsupportedValue.Wrap([]error{errors.New("rent amount can't be lower than 0")})})
	}
	if currentYear < 1 {
		err = err.AddToWrapped([]error{rent.ErrUnsupportedValue.Wrap([]error{errors.New("current year can't be lower than 1")})})
	}

	if len(err.UnWrap()) > 0 {
		return 0, err
	}

	rentIncrease := (rentAmount * (float32(currentYear-1) * (increaseRate / 100)))

	return (rentAmount + rentIncrease) * float32(s.configuration.MonthsInAYear), nil
}

func NewService(logger log.Logger, monthsInAYear int) rent.Service {
	return &rentService{
		logger, rent.Configuration{MonthsInAYear: monthsInAYear},
	}
}
