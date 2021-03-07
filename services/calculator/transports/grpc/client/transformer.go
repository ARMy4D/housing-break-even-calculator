package grpc

import (
	"context"
	"errors"

	"github.com/army4d/housing-break-even-calculator/models/pb"
	"github.com/army4d/housing-break-even-calculator/services/calculator"

	calculatorRequestResponses "github.com/army4d/housing-break-even-calculator/services/calculator/endpoints/request_responses"
)

func EncodeCalculateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(calculatorRequestResponses.CalculateRequest)
	if !ok {
		return nil, calculator.ErrTransformer.Wrap([]error{errors.New("encoding calculate grpc request")})
	}
	return &pb.CalculatorRequest{
		House: &pb.CalculatorRequest_HouseSetting{
			Price:                   req.House.Price,
			PropertyTaxRate:         req.House.PropertyTaxRate,
			PropertyTransferTaxRate: req.House.PropertyTransferTaxRate,
			YearsExpectedToReside:   int32(req.House.YearsExpectedToReside),
		},
		Rent: &pb.CalculatorRequest_RentSetting{
			Rent:                   req.Rent.Rent,
			YearlyRentIncreaseRate: req.Rent.YearlyRentIncreaseRate,
		},
		Mortgage: &pb.CalculatorRequest_MortgageSetting{
			DownPayment: req.Mortgage.DownPayment,
			IntrestRate: req.Mortgage.IntrestRate,
			Term:        int32(req.Mortgage.Term),
		},
	}, nil
}

func convertPBPaymentTypeToPaymentType(paymentType pb.PaymentType) calculator.PaymentType {
	switch paymentType {
	case pb.PaymentType_MORTGAGE:
		return calculator.PAYMENT_TYPE_MORTGAGE
	case pb.PaymentType_RENT:
		return calculator.PAYMENT_TYPE_RENT
	case pb.PaymentType_TOSS_UP:
		return calculator.PAYMENT_TYPE_TOSS_UP
	}
	return calculator.PaymentType("UKNOWN")
}

func DecodeCalculateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*pb.CalculatorResponse)
	if !ok {
		return nil, calculator.ErrTransformer.Wrap([]error{errors.New("decoding calculate grpc response")})
	}
	return calculatorRequestResponses.CalculateResponse{
		BreakEven: calculator.BreakEven{
			BreakEvenYearIntrest:                     int(resp.BreakEvenYearIntrest),
			BreakEvenYearOverall:                     int(resp.BreakEvenYearOverall),
			CostOfRentOverResidancePeriod:            resp.RentCostOverResidancePeriod,
			CostOfMortgageIntrestOverResidancePeriod: resp.MortgageIntrestCostOverResidancePeriod,
			CostOfMortgageOverResidancePeriod:        resp.MortgageCostOverResidancePeriod,
			BestPaymentTypeIntrest:                   convertPBPaymentTypeToPaymentType(resp.BestPaymentTypeIntrestOnly),
			BestPaymentTypeOverall:                   convertPBPaymentTypeToPaymentType(resp.BestPaymentTypeOverall),
		},
	}, nil
}
