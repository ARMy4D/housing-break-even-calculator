package calculator

type RentSetting struct {
	Rent                   float32
	YearlyRentIncreaseRate float32
}

type MortgageSetting struct {
	DownPayment float32
	IntrestRate float32
	Term        int
}

type HouseSetting struct {
	Price                   float32
	PropertyTaxRate         float32
	PropertyTransferTaxRate float32
	YearsExpectedToReside   int
}

type PaymentType string

const (
	PAYMENT_TYPE_RENT     PaymentType = "RENT"
	PAYMENT_TYPE_MORTGAGE PaymentType = "MORTGAGE"
	PAYMENT_TYPE_TOSS_UP  PaymentType = "TOSS_UP"
)

type BreakEven struct {
	BreakEvenYearIntrest                     int
	BreakEvenYearOverall                     int
	CostOfRentOverResidancePeriod            float32
	CostOfMortgageIntrestOverResidancePeriod float32
	CostOfMortgageOverResidancePeriod        float32
	BestPaymentTypeIntrest                   PaymentType
	BestPaymentTypeOverall                   PaymentType
}
