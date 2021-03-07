package rent

import (
	"context"
	"testing"

	"github.com/army4d/housing-break-even-calculator/services/rent"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

func makeNewRentService() rentService {
	return rentService{
		logger: log.NewNopLogger(),
		configuration: rent.Configuration{
			MonthsInAYear: 12,
		},
	}
}

type yearlyRentTestInput struct {
	rentAmount   float32
	currentYear  int
	increaseRate float32
}
type yearlyRentTestWant struct {
	totalPayment float32
	err          rent.Errorer
}

type yearlyRentTest struct {
	input yearlyRentTestInput
	want  yearlyRentTestWant
}

func TestCalculateYearlyRentConstraints(t *testing.T) {
	service := makeNewRentService()
	ctx := context.Background()

	tests := []yearlyRentTest{
		{
			input: yearlyRentTestInput{
				rentAmount:   -25000,
				currentYear:  1,
				increaseRate: 2,
			},
			want: yearlyRentTestWant{
				totalPayment: 0,
				err:          rent.ErrCalculation.Wrap([]error{rent.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyRentTestInput{
				rentAmount:   0,
				currentYear:  1,
				increaseRate: 2,
			},
			want: yearlyRentTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyRentTestInput{
				rentAmount:   1000,
				currentYear:  0,
				increaseRate: 2,
			},
			want: yearlyRentTestWant{
				totalPayment: 0,
				err:          rent.ErrCalculation.Wrap([]error{rent.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyRentTestInput{
				rentAmount:   1000,
				currentYear:  -1,
				increaseRate: 2,
			},
			want: yearlyRentTestWant{
				totalPayment: 0,
				err:          rent.ErrCalculation.Wrap([]error{rent.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyRentTestInput{
				rentAmount:   -1000,
				currentYear:  0,
				increaseRate: 2,
			},
			want: yearlyRentTestWant{
				totalPayment: 0,
				err:          rent.ErrCalculation.Wrap([]error{rent.ErrUnsupportedValue, rent.ErrUnsupportedValue}),
			},
		},
	}

	for _, currentTest := range tests {
		resp, err := service.CalculateYearlyRent(ctx, currentTest.input.rentAmount, currentTest.input.currentYear, currentTest.input.increaseRate)
		assert.Equal(t, currentTest.want.totalPayment, resp)
		if currentTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, currentTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(currentTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}

func TestCalculateYearlyRent(t *testing.T) {
	service := makeNewRentService()
	ctx := context.Background()

	tests := []yearlyRentTest{
		{
			input: yearlyRentTestInput{
				rentAmount:   0,
				currentYear:  2,
				increaseRate: 5,
			},
			want: yearlyRentTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyRentTestInput{
				rentAmount:   1000,
				currentYear:  10,
				increaseRate: 0,
			},
			want: yearlyRentTestWant{
				totalPayment: 12000,
				err:          nil,
			},
		},
		{
			input: yearlyRentTestInput{
				rentAmount:   1000,
				currentYear:  2,
				increaseRate: 5,
			},
			want: yearlyRentTestWant{
				totalPayment: 12600,
				err:          nil,
			},
		},
		{
			input: yearlyRentTestInput{
				rentAmount:   1000,
				currentYear:  1,
				increaseRate: 5,
			},
			want: yearlyRentTestWant{
				totalPayment: 12000,
				err:          nil,
			},
		},
	}

	for _, currentTest := range tests {
		resp, err := service.CalculateYearlyRent(ctx, currentTest.input.rentAmount, currentTest.input.currentYear, currentTest.input.increaseRate)
		assert.Equal(t, currentTest.want.totalPayment, resp)
		if currentTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, currentTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(currentTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}
