package calculator

import (
	"context"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"

	"github.com/army4d/housing-break-even-calculator/app"
	"github.com/army4d/housing-break-even-calculator/services/calculator"
	calculatorEndpoints "github.com/army4d/housing-break-even-calculator/services/calculator/endpoints"
	calculatorRequestResponses "github.com/army4d/housing-break-even-calculator/services/calculator/endpoints/request_responses"
)

type calcService struct {
	logger log.Logger

	configuration calculator.Configuration

	CalculateEndpoint endpoint.Endpoint
}

func (s calcService) Calculate(ctx context.Context, rent calculator.RentSetting, mortgage calculator.MortgageSetting, house calculator.HouseSetting) (calculator.BreakEven, calculator.Errorer) {

	// you do client side checking here

	if response, err := s.CalculateEndpoint(ctx, calculatorRequestResponses.CalculateRequest{
		Rent:     rent,
		Mortgage: mortgage,
		House:    house,
	}); err == nil {
		if resp, ok := response.(calculatorRequestResponses.CalculateResponse); ok {
			return calculator.BreakEven{
				BreakEvenYearIntrest:                     resp.BreakEvenYearIntrest,
				BreakEvenYearOverall:                     resp.BreakEvenYearOverall,
				CostOfRentOverResidancePeriod:            resp.CostOfRentOverResidancePeriod,
				CostOfMortgageOverResidancePeriod:        resp.CostOfMortgageOverResidancePeriod,
				CostOfMortgageIntrestOverResidancePeriod: resp.CostOfMortgageIntrestOverResidancePeriod,
				BestPaymentTypeIntrest:                   resp.BestPaymentTypeIntrest,
				BestPaymentTypeOverall:                   resp.BestPaymentTypeOverall,
			}, nil
		}
	} else {
		return calculator.BreakEven{}, calculator.ErrClient.Wrap([]error{err})
	}
	return calculator.BreakEven{}, calculator.ErrClient
}

func NewService(logger log.Logger, conn *grpc.ClientConn, rateLimitDuration time.Duration, language app.ErrorTranslation) calculator.Service {

	var calculateEndpoint endpoint.Endpoint
	{
		// you could add client options to the endpoint here
		calculateEndpoint = calculatorEndpoints.MakeCalculateClientEndpoint(conn, kitgrpc.ClientBefore(kitgrpc.SetRequestHeader("lang", string(language))))
		calculateEndpoint = calculatorEndpoints.LoggingMiddleware(log.With(logger, "method", "calculate"))(calculateEndpoint)
		calculateEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(rateLimitDuration), 1))(calculateEndpoint)
		calculateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Calculate",
			Timeout: 30 * time.Second,
		}))(calculateEndpoint)
	}

	return &calcService{
		logger, calculator.Configuration{}, calculateEndpoint,
	}
}
