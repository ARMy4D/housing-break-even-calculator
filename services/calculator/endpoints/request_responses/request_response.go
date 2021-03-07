package calculator

import "github.com/army4d/housing-break-even-calculator/services/calculator"

type CalculateRequest struct {
	Rent     calculator.RentSetting
	Mortgage calculator.MortgageSetting
	House    calculator.HouseSetting
}

type CalculateResponse struct {
	calculator.BreakEven
}
