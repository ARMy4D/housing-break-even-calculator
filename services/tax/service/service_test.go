package tax

import (
	"context"
	"testing"

	"github.com/army4d/housing-break-even-calculator/services/tax"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

func makeNewTaxService() taxService {
	return taxService{
		logger:        log.NewNopLogger(),
		configuration: tax.Configuration{},
	}
}

type yearlyTaxesTestInput struct {
	housePrice  float32
	currentYear int
	taxes       tax.TaxSetting
}
type yearlyTaxesTestWant struct {
	totalPayment float32
	err          tax.Errorer
}

type yearlyTaxesTest struct {
	input yearlyTaxesTestInput
	want  yearlyTaxesTestWant
}

func TestCalculateYearlyTaxesConstraints(t *testing.T) {
	service := makeNewTaxService()
	ctx := context.Background()

	tests := []yearlyTaxesTest{
		{
			input: yearlyTaxesTestInput{
				housePrice:  -25000,
				currentYear: 1,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         0,
					PropertyTransferTaxRate: 0,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 0,
				err:          tax.ErrCalculation.Wrap([]error{tax.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyTaxesTestInput{
				housePrice:  0,
				currentYear: 1,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         5,
					PropertyTransferTaxRate: 5,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyTaxesTestInput{
				housePrice:  1000,
				currentYear: 0,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         0,
					PropertyTransferTaxRate: 0,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 0,
				err:          tax.ErrCalculation.Wrap([]error{tax.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyTaxesTestInput{
				housePrice:  1000,
				currentYear: -1,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         0,
					PropertyTransferTaxRate: 0,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 0,
				err:          tax.ErrCalculation.Wrap([]error{tax.ErrUnsupportedValue}),
			},
		},
		{
			input: yearlyTaxesTestInput{
				housePrice:  -1000,
				currentYear: 0,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         0,
					PropertyTransferTaxRate: 0,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 0,
				err:          tax.ErrCalculation.Wrap([]error{tax.ErrUnsupportedValue, tax.ErrUnsupportedValue}),
			},
		},
	}

	for _, curtaxTest := range tests {
		resp, err := service.CalculateYearlyTaxes(ctx, curtaxTest.input.housePrice, curtaxTest.input.currentYear, curtaxTest.input.taxes)
		assert.Equal(t, curtaxTest.want.totalPayment, resp)
		if curtaxTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, curtaxTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(curtaxTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}

func TestCalculateYearlyTax(t *testing.T) {
	service := makeNewTaxService()
	ctx := context.Background()

	tests := []yearlyTaxesTest{
		{
			input: yearlyTaxesTestInput{
				housePrice:  0,
				currentYear: 2,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         6,
					PropertyTransferTaxRate: 1,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
		{
			input: yearlyTaxesTestInput{
				housePrice:  10000,
				currentYear: 10,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         5,
					PropertyTransferTaxRate: 2,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 500,
				err:          nil,
			},
		},
		{
			input: yearlyTaxesTestInput{
				housePrice:  1000,
				currentYear: 1,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         5,
					PropertyTransferTaxRate: 2,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 70,
				err:          nil,
			},
		},
		{
			input: yearlyTaxesTestInput{
				housePrice:  1000,
				currentYear: 1,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         0,
					PropertyTransferTaxRate: 5,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 50,
				err:          nil,
			},
		},
		{
			input: yearlyTaxesTestInput{
				housePrice:  1000,
				currentYear: 1,
				taxes: tax.TaxSetting{
					PropertyTaxRate:         0,
					PropertyTransferTaxRate: 0,
				},
			},
			want: yearlyTaxesTestWant{
				totalPayment: 0,
				err:          nil,
			},
		},
	}

	for _, curtaxTest := range tests {
		resp, err := service.CalculateYearlyTaxes(ctx, curtaxTest.input.housePrice, curtaxTest.input.currentYear, curtaxTest.input.taxes)
		assert.Equal(t, curtaxTest.want.totalPayment, resp)
		if curtaxTest.want.err == nil {
			assert.NoError(t, err)
		} else {
			assert.Equal(t, curtaxTest.want.err.ErrCode(), err.ErrCode())
			assert.Equal(t, len(curtaxTest.want.err.UnWrap()), len(err.UnWrap()))
		}
	}
}
