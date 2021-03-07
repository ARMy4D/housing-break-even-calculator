package grpc

import (
	"context"
	"errors"

	"github.com/army4d/housing-break-even-calculator/models/pb"
	"github.com/army4d/housing-break-even-calculator/services/calculator"
	calculatorRequestResponses "github.com/army4d/housing-break-even-calculator/services/calculator/endpoints/request_responses"
)

func decodeCalculateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*pb.CalculatorRequest)
	if !ok {
		return nil, calculator.ErrTransformer.Wrap([]error{errors.New("decoding calculate grpc request")})
	}
	return calculatorRequestResponses.CalculateRequest{
		Rent:     calculator.RentSetting{Rent: req.Rent.Rent, YearlyRentIncreaseRate: req.Rent.YearlyRentIncreaseRate},
		Mortgage: calculator.MortgageSetting{DownPayment: req.Mortgage.DownPayment, IntrestRate: req.Mortgage.IntrestRate, Term: int(req.Mortgage.Term)},
		House:    calculator.HouseSetting{Price: req.House.Price, PropertyTaxRate: req.House.PropertyTaxRate, PropertyTransferTaxRate: req.House.PropertyTransferTaxRate, YearsExpectedToReside: int(req.House.YearsExpectedToReside)},
	}, nil
}

func convertPaymentTypeToPBPaymentType(paymentType calculator.PaymentType) pb.PaymentType {
	switch paymentType {
	case calculator.PAYMENT_TYPE_MORTGAGE:
		return pb.PaymentType_MORTGAGE
	case calculator.PAYMENT_TYPE_RENT:
		return pb.PaymentType_RENT
	case calculator.PAYMENT_TYPE_TOSS_UP:
		return pb.PaymentType_TOSS_UP
	}
	return pb.PaymentType(-1)
}

func encodeCalculateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(calculatorRequestResponses.CalculateResponse)
	if !ok {
		return nil, calculator.ErrTransformer.Wrap([]error{errors.New("encoding calculate grpc response")})
	}
	return &pb.CalculatorResponse{
		BreakEvenYearIntrest:                   int32(resp.BreakEvenYearIntrest),
		BreakEvenYearOverall:                   int32(resp.BreakEvenYearOverall),
		MortgageIntrestCostOverResidancePeriod: resp.CostOfMortgageIntrestOverResidancePeriod,
		MortgageCostOverResidancePeriod:        resp.CostOfMortgageOverResidancePeriod,
		RentCostOverResidancePeriod:            resp.CostOfRentOverResidancePeriod,
		BestPaymentTypeIntrestOnly:             convertPaymentTypeToPBPaymentType(resp.BestPaymentTypeIntrest),
		BestPaymentTypeOverall:                 convertPaymentTypeToPBPaymentType(resp.BestPaymentTypeOverall),
	}, nil
}
