package calculator

import (
	"context"

	"github.com/army4d/housing-break-even-calculator/services/calculator"
	"github.com/army4d/housing-break-even-calculator/services/mortgage"
	"github.com/army4d/housing-break-even-calculator/services/rent"
	"github.com/army4d/housing-break-even-calculator/services/tax"
	"github.com/go-kit/kit/log"
)

type calcService struct {
	logger log.Logger

	configuration calculator.Configuration

	mSvc mortgage.Service
	tSvc tax.Service
	rSvc rent.Service
}

func (s calcService) Calculate(ctx context.Context, rent calculator.RentSetting, mortgage calculator.MortgageSetting, house calculator.HouseSetting) (calculator.BreakEven, calculator.Errorer) {
	var rentCost, mortgageCost, repaymentAmount float32

	breakEven := calculator.BreakEven{
		BestPaymentTypeIntrest: calculator.PAYMENT_TYPE_RENT,
		BestPaymentTypeOverall: calculator.PAYMENT_TYPE_RENT,
	}

	for year := 1; year <= house.YearsExpectedToReside; year++ {
		if currentYearCost, err := s.rSvc.CalculateYearlyRent(ctx, rent.Rent, year, rent.YearlyRentIncreaseRate); err == nil {
			rentCost += currentYearCost
		} else {
			return calculator.BreakEven{}, calculator.ErrParameter.Wrap([]error{err})
		}

		if currentYearCost, err := s.mSvc.CalculateYearlyMortgageIntrest(ctx, house.Price-mortgage.DownPayment, year, mortgage.Term, mortgage.IntrestRate); err == nil {
			mortgageCost += currentYearCost
		} else {
			return calculator.BreakEven{}, calculator.ErrParameter.Wrap([]error{err})
		}

		if currentYearCost, err := s.mSvc.CalculateYearlyMortgageRepayment(ctx, house.Price-mortgage.DownPayment, year, mortgage.Term); err == nil {
			repaymentAmount += currentYearCost
		} else {
			return calculator.BreakEven{}, calculator.ErrParameter.Wrap([]error{err})
		}

		if currentYearCost, err := s.tSvc.CalculateYearlyTaxes(ctx, house.Price, year, tax.TaxSetting{PropertyTaxRate: house.PropertyTaxRate, PropertyTransferTaxRate: house.PropertyTransferTaxRate}); err == nil {
			mortgageCost += currentYearCost
		} else {
			return calculator.BreakEven{}, calculator.ErrParameter.Wrap([]error{err})
		}

		if mortgageCost < rentCost && breakEven.BreakEvenYearIntrest == 0 {
			breakEven.BreakEvenYearIntrest = year
			breakEven.BestPaymentTypeIntrest = calculator.PAYMENT_TYPE_MORTGAGE
		}

		if mortgageCost+repaymentAmount < rentCost && breakEven.BreakEvenYearOverall == 0 {
			breakEven.BreakEvenYearOverall = year
			breakEven.BestPaymentTypeOverall = calculator.PAYMENT_TYPE_MORTGAGE
		}
	}

	if mortgageCost == rentCost {
		breakEven.BestPaymentTypeIntrest = calculator.PAYMENT_TYPE_TOSS_UP
	}
	if mortgageCost+repaymentAmount == rentCost {
		breakEven.BestPaymentTypeOverall = calculator.PAYMENT_TYPE_TOSS_UP
	}

	breakEven.CostOfMortgageIntrestOverResidancePeriod = mortgageCost
	breakEven.CostOfMortgageOverResidancePeriod = mortgageCost + repaymentAmount
	breakEven.CostOfRentOverResidancePeriod = rentCost

	return breakEven, nil
}

func NewService(logger log.Logger, mSvc mortgage.Service, tSvc tax.Service, rSvc rent.Service) calculator.Service {
	return &calcService{
		logger, calculator.Configuration{},
		mSvc,
		tSvc,
		rSvc,
	}
}
