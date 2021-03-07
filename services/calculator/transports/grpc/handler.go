package grpc

import (
	"context"
	"time"

	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/army4d/housing-break-even-calculator/models/pb"
	"github.com/army4d/housing-break-even-calculator/services/calculator"
	calculatorEndpoints "github.com/army4d/housing-break-even-calculator/services/calculator/endpoints"
)

type server struct {
	pb.UnimplementedCalculatorServer
	calculate kitgrpc.Handler
}

func (s *server) Calculate(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	_, resp, err := s.calculate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CalculatorResponse), nil
}

func MakeHandler(s calculator.Service, logger log.Logger, rateLimitDuration time.Duration) pb.CalculatorServer {
	opts := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorHandler(calculator.NewErrorHandlerGRPC(logger)),
	}

	calculateEndpoint := calculatorEndpoints.MakeCalculateEndpoint(s)
	{
		calculateEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(rateLimitDuration), 1))(calculateEndpoint)
		calculateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(calculateEndpoint)
		calculateEndpoint = calculatorEndpoints.ErrorTranslateMiddleware(calculateEndpoint)
		calculateEndpoint = calculatorEndpoints.LoggingMiddleware(log.With(logger, "method", "calculate"))(calculateEndpoint)
	}

	return &server{
		calculate: kitgrpc.NewServer(calculateEndpoint, calculator.ErrorHandlerMiddlewareGRPC(decodeCalculateRequest), calculator.ErrorHandlerMiddlewareGRPC(encodeCalculateResponse), opts...),
	}
}
