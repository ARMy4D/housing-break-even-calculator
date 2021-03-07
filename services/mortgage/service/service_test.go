package mortgage

import (
	"context"
	"testing"

	"github.com/army4d/housing-break-even-calculator/services/mortgage"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

func makeNewMortService() mortService {
	return mortService{
		logger:        log.NewNopLogger(),
		configuration: mortgage.Configuration{},
	}
}

type mortgagePaidAmountTestInput struct {
	mortgageAmount float32
	currentYear    int
	termInYears    int
}
type mortgagePaidAmountTestWant struct {
	totalPayment float32
	err          mortgage.Errorer
}

type mortgagePaidAmountTest struct {
	input mortgagePaidAmountTestInput
	want  mortgagePaidAmountTestWant
}

func TestCalculateMortgagePaidAmountConstraints(t *testing.T) {
	service := makeNewMortService()

	tests := []mortgagePaidAmountTest{
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: -25000,
				currentYear:    2,
				termInYears:    5,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 0,
				currentYear:    1,
				termInYears:    5,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 1000,
				currentYear:    -5,
				termInYears:    5,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 1000,
				currentYear:    0,
				termInYears:    5,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 1000,
				currentYear:    5,
				termInYears:    -5,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 1000,
				currentYear:    5,
				termInYears:    0,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: -1000,
				currentYear:    -1,
				termInYears:    -1,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: -1000,
				currentYear:    1,
				termInYears:    -1,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 1000,
				currentYear:    -1,
				termInYears:    -1,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: -1000,
				currentYear:    -1,
				termInYears:    1,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
	}

	for _, currentTest := range tests {
		resp, err := service.CalculateMortgagePaidAmount(currentTest.input.mortgageAmount, currentTest.input.currentYear, currentTest.input.termInYears)
		assert.Equal(t, currentTest.want.totalPayment, resp)
		if currentTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, currentTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(currentTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}

func TestCalculateMortgagePaidAmount(t *testing.T) {
	service := makeNewMortService()

	tests := []mortgagePaidAmountTest{
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 0,
				currentYear:    2,
				termInYears:    5,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 1000,
				currentYear:    10,
				termInYears:    2,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 1000,
				err:          nil,
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 1000,
				currentYear:    2,
				termInYears:    2,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 500,
				err:          nil,
			},
		},
		{
			input: mortgagePaidAmountTestInput{
				mortgageAmount: 1000,
				currentYear:    1,
				termInYears:    2,
			},
			want: mortgagePaidAmountTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
	}

	for _, currentTest := range tests {
		resp, err := service.CalculateMortgagePaidAmount(currentTest.input.mortgageAmount, currentTest.input.currentYear, currentTest.input.termInYears)
		assert.Equal(t, currentTest.want.totalPayment, resp)
		if currentTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, currentTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(currentTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}

}

type yearlyMortgageRepaymentTestInput struct {
	mortgageAmount float32
	currentYear    int
	termInYears    int
}
type yearlyMortgageRepaymentTestWant struct {
	totalPayment float32
	err          mortgage.Errorer
}

type yearlyMortgageRepaymentTest struct {
	input yearlyMortgageRepaymentTestInput
	want  yearlyMortgageRepaymentTestWant
}

func TestCalculateYearlyMortgageRepaymentConstraints(t *testing.T) {
	service := makeNewMortService()
	ctx := context.Background()

	tests := []yearlyMortgageRepaymentTest{
		{
			input: yearlyMortgageRepaymentTestInput{
				mortgageAmount: -25000,
				currentYear:    0,
				termInYears:    5,
			},
			want: yearlyMortgageRepaymentTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageRepaymentTestInput{
				mortgageAmount: 0,
				currentYear:    0,
				termInYears:    5,
			},
			want: yearlyMortgageRepaymentTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyMortgageRepaymentTestInput{
				mortgageAmount: 1000,
				currentYear:    0,
				termInYears:    -5,
			},
			want: yearlyMortgageRepaymentTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageRepaymentTestInput{
				mortgageAmount: 1000,
				currentYear:    0,
				termInYears:    0,
			},
			want: yearlyMortgageRepaymentTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageRepaymentTestInput{
				mortgageAmount: -1000,
				currentYear:    0,
				termInYears:    -1,
			},
			want: yearlyMortgageRepaymentTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
	}

	for _, currentTest := range tests {
		resp, err := service.CalculateYearlyMortgageRepayment(ctx, currentTest.input.mortgageAmount, currentTest.input.currentYear, currentTest.input.termInYears)
		assert.Equal(t, currentTest.want.totalPayment, resp)
		if currentTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, currentTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(currentTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}

func TestCalculateYearlyMortgageRepayment(t *testing.T) {
	service := makeNewMortService()
	ctx := context.Background()

	tests := []yearlyMortgageRepaymentTest{
		{
			input: yearlyMortgageRepaymentTestInput{
				mortgageAmount: 0,
				currentYear:    0,
				termInYears:    5,
			},
			want: yearlyMortgageRepaymentTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyMortgageRepaymentTestInput{
				mortgageAmount: 1000,
				currentYear:    0,
				termInYears:    2,
			},
			want: yearlyMortgageRepaymentTestWant{
				totalPayment: 500,
				err:          nil,
			},
		},
		{
			input: yearlyMortgageRepaymentTestInput{
				mortgageAmount: 1000,
				currentYear:    12,
				termInYears:    2,
			},
			want: yearlyMortgageRepaymentTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
	}

	for _, currentTest := range tests {
		resp, err := service.CalculateYearlyMortgageRepayment(ctx, currentTest.input.mortgageAmount, currentTest.input.currentYear, currentTest.input.termInYears)
		assert.Equal(t, currentTest.want.totalPayment, resp)
		if currentTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, currentTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(currentTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}

type yearlyMortgageIntrestTestInput struct {
	mortgageAmount float32
	currentYear    int
	termInYears    int
	intrestRate    float32
}
type yearlyMortgageIntrestTestWant struct {
	totalPayment float32
	err          mortgage.Errorer
}

type yearlyMortgageIntrestTest struct {
	input yearlyMortgageIntrestTestInput
	want  yearlyMortgageIntrestTestWant
}

func TestCalculateYearlyMortgageIntrestConstraints(t *testing.T) {
	service := makeNewMortService()
	ctx := context.Background()

	tests := []yearlyMortgageIntrestTest{
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: -25000,
				currentYear:    2,
				termInYears:    5,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 0,
				currentYear:    1,
				termInYears:    5,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 1000,
				currentYear:    -5,
				termInYears:    5,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 1000,
				currentYear:    0,
				termInYears:    5,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 1000,
				currentYear:    5,
				termInYears:    -5,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 1000,
				currentYear:    5,
				termInYears:    0,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: -1000,
				currentYear:    -1,
				termInYears:    -1,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: -1000,
				currentYear:    1,
				termInYears:    -1,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 1000,
				currentYear:    -1,
				termInYears:    -1,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: -1000,
				currentYear:    -1,
				termInYears:    1,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          mortgage.ErrCalculation.Wrap([]error{mortgage.ErrUnsupportedValue, mortgage.ErrUnsupportedValue}),
			},
		},
	}

	for _, currentTest := range tests {
		resp, err := service.CalculateYearlyMortgageIntrest(ctx, currentTest.input.mortgageAmount, currentTest.input.currentYear, currentTest.input.termInYears, currentTest.input.intrestRate)
		assert.Equal(t, currentTest.want.totalPayment, resp)
		if currentTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, currentTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(currentTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}

func TestCalculateYearlyMortgageIntrest(t *testing.T) {
	service := makeNewMortService()
	ctx := context.Background()

	tests := []yearlyMortgageIntrestTest{
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 0,
				currentYear:    2,
				termInYears:    5,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 2000,
				currentYear:    6,
				termInYears:    5,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 2000,
				currentYear:    2,
				termInYears:    2,
				intrestRate:    2,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 20,
				err:          nil,
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 2000,
				currentYear:    5,
				termInYears:    5,
				intrestRate:    5,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 20,
				err:          nil,
			},
		},
		{
			input: yearlyMortgageIntrestTestInput{
				mortgageAmount: 2000,
				currentYear:    1,
				termInYears:    5,
				intrestRate:    5,
			},
			want: yearlyMortgageIntrestTestWant{
				totalPayment: 100,
				err:          nil,
			},
		},
	}

	for _, currentTest := range tests {
		resp, err := service.CalculateYearlyMortgageIntrest(ctx, currentTest.input.mortgageAmount, currentTest.input.currentYear, currentTest.input.termInYears, currentTest.input.intrestRate)
		assert.Equal(t, currentTest.want.totalPayment, resp)
		if currentTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, currentTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(currentTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}
