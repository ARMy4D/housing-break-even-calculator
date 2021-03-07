package calculator

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/army4d/housing-break-even-calculator/app"
)

func ErrorHandlerMiddlewareGRPC(next func(context.Context, interface{}) (interface{}, error)) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		resp, err := next(ctx, input)
		if err != nil {
			if e, ok := err.(Errorer); ok {
				if md, ok := metadata.FromIncomingContext(ctx); ok && md.Len() > 0 {
					lang := md.Get("lang")
					if len(lang) == 1 {
						switch lang[0] {
						case string(app.ERR_EN):
							err = e.SetTranslation(app.ERR_EN)
						case string(app.ERR_NL):
							err = e.SetTranslation(app.ERR_NL)
						}
					}
				}
				k, v := kitgrpc.EncodeKeyValue("error_code", fmt.Sprintf("%v", e.ErrCode()))
				metadata.AppendToOutgoingContext(ctx, k, v)
				switch e.ErrCode() {
				case ErrConstraintReached.ErrCode(), ErrMaxMortgageTerm.ErrCode(), ErrMinMortgageTerm.ErrCode(), ErrMinReside.ErrCode(), ErrParameter.ErrCode():
					err = status.Error(codes.InvalidArgument, err.Error())
				case ErrTransformer.ErrCode():
					err = status.Error(codes.Internal, err.Error())
				}
			}
		}
		return resp, err
	}
}

type errorHandler struct {
	logger log.Logger
}

func (eh errorHandler) Handle(ctx context.Context, err error) {
	level.Error(eh.logger).Log("transport:", "grpc", "service", "calculator")
}

func NewErrorHandlerGRPC(logger log.Logger) transport.ErrorHandler {
	return errorHandler{
		logger: logger,
	}
}
