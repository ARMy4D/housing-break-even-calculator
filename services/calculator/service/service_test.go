package calculator

import (
	"context"
	"testing"

	"github.com/army4d/housing-break-even-calculator/services/calculator"
	"github.com/army4d/housing-break-even-calculator/services/mortgage"
	"github.com/army4d/housing-break-even-calculator/services/rent"
	"github.com/army4d/housing-break-even-calculator/services/tax"
	"github.com/go-kit/kit/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mSvcMock struct {
	mock.Mock
}

func (m *mSvcMock) CalculateYearlyMortgageIntrest(ctx context.Context, mortgageAmount float32, currentYear int, termInYears int, intrestRate float32) (totalPayment float32, err mortgage.Errorer) {
	args := m.Called(ctx, mortgageAmount, currentYear, termInYears, intrestRate)
	if err, ok := args.Get(1).(mortgage.Errorer); ok {
		return args.Get(0).(float32), err
	}
	return args.Get(0).(float32), nil
}
func (m *mSvcMock) CalculateYearlyMortgageRepayment(ctx context.Context, mortgageAmount float32, currentYear int, termInYears int) (totalPayment float32, err mortgage.Errorer) {
	args := m.Called(ctx, mortgageAmount, currentYear, termInYears)
	if err, ok := args.Get(1).(mortgage.Errorer); ok {
		return args.Get(0).(float32), err
	}
	return args.Get(0).(float32), nil
}

type rSvcMock struct {
	mock.Mock
}

func (m *rSvcMock) CalculateYearlyRent(ctx context.Context, rentAmount float32, currentYear int, increaseRate float32) (totalPayment float32, err rent.Errorer) {
	args := m.Called(ctx, rentAmount, currentYear, increaseRate)
	if err, ok := args.Get(1).(rent.Errorer); ok {
		return args.Get(0).(float32), err
	}
	return args.Get(0).(float32), nil
}

type tSvcMock struct {
	mock.Mock
}

func (m *tSvcMock) CalculateYearlyTaxes(ctx context.Context, housePrice float32, currentYear int, taxes tax.TaxSetting) (totalPayment float32, err tax.Errorer) {
	args := m.Called(ctx, housePrice, currentYear, taxes)
	if err, ok := args.Get(1).(tax.Errorer); ok {
		return args.Get(0).(float32), err
	}
	return args.Get(0).(float32), nil
}
func TestCalculateNoResidanceExpectation(t *testing.T) {

	msvc := new(mSvcMock)
	rsvc := new(rSvcMock)
	tsvc := new(tSvcMock)

	service := NewService(log.NewNopLogger(), msvc, tsvc, rsvc)

	resp, err := service.Calculate(
		context.Background(),
		calculator.RentSetting{
			Rent:                   2000,
			YearlyRentIncreaseRate: 2,
		},
		calculator.MortgageSetting{
			DownPayment: 2000,
			IntrestRate: 2,
			Term:        3,
		},
		calculator.HouseSetting{
			Price:                   5000,
			PropertyTaxRate:         2,
			PropertyTransferTaxRate: 1,
			YearsExpectedToReside:   0,
		},
	)

	assert.Equal(t, calculator.BreakEven{
		BestPaymentTypeIntrest:                   calculator.PAYMENT_TYPE_TOSS_UP,
		BestPaymentTypeOverall:                   calculator.PAYMENT_TYPE_TOSS_UP,
		BreakEvenYearIntrest:                     0,
		BreakEvenYearOverall:                     0,
		CostOfRentOverResidancePeriod:            0,
		CostOfMortgageIntrestOverResidancePeriod: 0,
		CostOfMortgageOverResidancePeriod:        0,
	}, resp)
	assert.NoError(t, err)

	msvc.AssertNotCalled(t, "CalculateYearlyMortgageIntrest")
	msvc.AssertNotCalled(t, "CalculateYearlyMortgageRepayment")

	rsvc.AssertNotCalled(t, "CalculateYearlyRent")

	tsvc.AssertNotCalled(t, "CalculateYearlyTaxes")
}

func TestCalculateRent(t *testing.T) {

	msvc := new(mSvcMock)

	msvc.On("CalculateYearlyMortgageIntrest", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(10), nil)
	msvc.On("CalculateYearlyMortgageRepayment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(20), nil)

	rsvc := new(rSvcMock)

	rsvc.On("CalculateYearlyRent", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(20), nil)

	tsvc := new(tSvcMock)

	tsvc.On("CalculateYearlyTaxes", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(20), nil)

	service := NewService(log.NewNopLogger(), msvc, tsvc, rsvc)

	resp, err := service.Calculate(
		context.Background(),
		calculator.RentSetting{
			Rent:                   2000,
			YearlyRentIncreaseRate: 2,
		},
		calculator.MortgageSetting{
			DownPayment: 2000,
			IntrestRate: 2,
			Term:        3,
		},
		calculator.HouseSetting{
			Price:                   5000,
			PropertyTaxRate:         2,
			PropertyTransferTaxRate: 1,
			YearsExpectedToReside:   5,
		},
	)

	assert.Equal(t, calculator.BreakEven{
		BestPaymentTypeIntrest:                   calculator.PAYMENT_TYPE_RENT,
		BestPaymentTypeOverall:                   calculator.PAYMENT_TYPE_RENT,
		BreakEvenYearIntrest:                     0,
		BreakEvenYearOverall:                     0,
		CostOfRentOverResidancePeriod:            100,
		CostOfMortgageIntrestOverResidancePeriod: 150,
		CostOfMortgageOverResidancePeriod:        250,
	}, resp)
	assert.NoError(t, err)

	msvc.AssertNumberOfCalls(t, "CalculateYearlyMortgageIntrest", 5)
	msvc.AssertNumberOfCalls(t, "CalculateYearlyMortgageRepayment", 5)

	rsvc.AssertNumberOfCalls(t, "CalculateYearlyRent", 5)

	tsvc.AssertNumberOfCalls(t, "CalculateYearlyTaxes", 5)

	msvc.AssertExpectations(t)
	rsvc.AssertExpectations(t)
	tsvc.AssertExpectations(t)
}

func TestCalculateMortgage(t *testing.T) {

	msvc := new(mSvcMock)

	msvc.On("CalculateYearlyMortgageIntrest", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(10), nil)
	msvc.On("CalculateYearlyMortgageRepayment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(20), nil)

	rsvc := new(rSvcMock)

	rsvc.On("CalculateYearlyRent", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(60), nil)

	tsvc := new(tSvcMock)

	tsvc.On("CalculateYearlyTaxes", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(20), nil)

	service := NewService(log.NewNopLogger(), msvc, tsvc, rsvc)

	resp, err := service.Calculate(
		context.Background(),
		calculator.RentSetting{
			Rent:                   2000,
			YearlyRentIncreaseRate: 2,
		},
		calculator.MortgageSetting{
			DownPayment: 2000,
			IntrestRate: 2,
			Term:        3,
		},
		calculator.HouseSetting{
			Price:                   5000,
			PropertyTaxRate:         2,
			PropertyTransferTaxRate: 1,
			YearsExpectedToReside:   5,
		},
	)

	assert.Equal(t, calculator.BreakEven{
		BestPaymentTypeIntrest:                   calculator.PAYMENT_TYPE_MORTGAGE,
		BestPaymentTypeOverall:                   calculator.PAYMENT_TYPE_MORTGAGE,
		BreakEvenYearIntrest:                     1,
		BreakEvenYearOverall:                     1,
		CostOfRentOverResidancePeriod:            300,
		CostOfMortgageIntrestOverResidancePeriod: 150,
		CostOfMortgageOverResidancePeriod:        250,
	}, resp)
	assert.NoError(t, err)

	msvc.AssertNumberOfCalls(t, "CalculateYearlyMortgageIntrest", 5)
	msvc.AssertNumberOfCalls(t, "CalculateYearlyMortgageRepayment", 5)

	rsvc.AssertNumberOfCalls(t, "CalculateYearlyRent", 5)

	tsvc.AssertNumberOfCalls(t, "CalculateYearlyTaxes", 5)

	msvc.AssertExpectations(t)
	rsvc.AssertExpectations(t)
	tsvc.AssertExpectations(t)
}

func TestCalculateMixed(t *testing.T) {

	msvc := new(mSvcMock)

	msvc.On("CalculateYearlyMortgageIntrest", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(10), nil)
	msvc.On("CalculateYearlyMortgageRepayment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(20), nil)

	rsvc := new(rSvcMock)

	rsvc.On("CalculateYearlyRent", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(50), nil)

	tsvc := new(tSvcMock)

	tsvc.On("CalculateYearlyTaxes", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float32(20), nil)

	service := NewService(log.NewNopLogger(), msvc, tsvc, rsvc)

	resp, err := service.Calculate(
		context.Background(),
		calculator.RentSetting{
			Rent:                   2000,
			YearlyRentIncreaseRate: 2,
		},
		calculator.MortgageSetting{
			DownPayment: 2000,
			IntrestRate: 2,
			Term:        3,
		},
		calculator.HouseSetting{
			Price:                   5000,
			PropertyTaxRate:         2,
			PropertyTransferTaxRate: 1,
			YearsExpectedToReside:   5,
		},
	)

	assert.Equal(t, calculator.BreakEven{
		BestPaymentTypeIntrest:                   calculator.PAYMENT_TYPE_MORTGAGE,
		BestPaymentTypeOverall:                   calculator.PAYMENT_TYPE_TOSS_UP,
		BreakEvenYearIntrest:                     1,
		BreakEvenYearOverall:                     0,
		CostOfRentOverResidancePeriod:            250,
		CostOfMortgageIntrestOverResidancePeriod: 150,
		CostOfMortgageOverResidancePeriod:        250,
	}, resp)
	assert.NoError(t, err)

	msvc.AssertNumberOfCalls(t, "CalculateYearlyMortgageIntrest", 5)
	msvc.AssertNumberOfCalls(t, "CalculateYearlyMortgageRepayment", 5)

	rsvc.AssertNumberOfCalls(t, "CalculateYearlyRent", 5)

	tsvc.AssertNumberOfCalls(t, "CalculateYearlyTaxes", 5)

	msvc.AssertExpectations(t)
	rsvc.AssertExpectations(t)
	tsvc.AssertExpectations(t)
}
