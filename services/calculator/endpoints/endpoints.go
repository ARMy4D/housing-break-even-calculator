package calculator

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/army4d/housing-break-even-calculator/models/pb"
	"github.com/army4d/housing-break-even-calculator/services/calculator"
	calculatorRequestResponses "github.com/army4d/housing-break-even-calculator/services/calculator/endpoints/request_responses"
	calculatorTrasformers "github.com/army4d/housing-break-even-calculator/services/calculator/transports/grpc/client"
)

func MakeCalculateEndpoint(s calculator.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(calculatorRequestResponses.CalculateRequest)
		resp, er := s.Calculate(ctx, req.Rent, req.Mortgage, req.House)
		if er != nil {
			return nil, er
		}
		return calculatorRequestResponses.CalculateResponse{resp}, nil
	}
}

func MakeCalculateClientEndpoint(conn *grpc.ClientConn, opts ...kitgrpc.ClientOption) endpoint.Endpoint {
	return kitgrpc.NewClient(
		conn,
		"Calculator",
		"Calculate",
		calculatorTrasformers.EncodeCalculateRequest,
		calculatorTrasformers.DecodeCalculateResponse,
		pb.CalculatorResponse{},
		opts...,
	).Endpoint()
}

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				if err != nil {
					level.Error(logger).Log("transport_error", err, "took", time.Since(begin))
				} else {
					level.Info(logger).Log("took", time.Since(begin))
				}
			}(time.Now())
			return next(ctx, request)

		}
	}
}

func ErrorTranslateMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return calculator.ErrorHandlerMiddlewareGRPC(next)(ctx, request)
	}
}
