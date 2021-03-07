package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"google.golang.org/grpc"

	"github.com/army4d/housing-break-even-calculator/models/pb"

	"github.com/army4d/housing-break-even-calculator/app"
	"github.com/go-kit/kit/log/level"

	"github.com/army4d/housing-break-even-calculator/services/rent"
	rentSvc "github.com/army4d/housing-break-even-calculator/services/rent/service"

	"github.com/army4d/housing-break-even-calculator/services/mortgage"
	mortgageSvc "github.com/army4d/housing-break-even-calculator/services/mortgage/service"

	"github.com/army4d/housing-break-even-calculator/services/tax"
	taxSvc "github.com/army4d/housing-break-even-calculator/services/tax/service"

	"github.com/army4d/housing-break-even-calculator/services/calculator"
	calculatorSvc "github.com/army4d/housing-break-even-calculator/services/calculator/service"
	calculatorGrpc "github.com/army4d/housing-break-even-calculator/services/calculator/transports/grpc"
)

const (
	defaultConfigName = "default"
	serviceName       = "calculator"
)

func main() {

	config, err := app.InitConfig(defaultConfigName)
	if err != nil {
		panic(fmt.Errorf("fatal error in config file: %v+", err))
	}
	logger := app.InitLogger(config.GetBool("debug"), serviceName)

	var rsvc rent.Service
	{
		rsvc = rentSvc.NewService(
			logger,
			config.GetInt("months_in_a_year"),
		)
	}

	var msvc mortgage.Service
	{
		msvc = mortgageSvc.NewService(
			logger,
		)
	}

	var tsvc tax.Service
	{
		tsvc = taxSvc.NewService(
			logger,
		)
	}

	var csvc calculator.Service
	{
		csvc = calculatorSvc.NewService(
			logger,
			msvc,
			tsvc,
			rsvc,
		)
		csvc = calculatorSvc.NewValidationMiddleware(logger,
			config.GetInt("min_year_to_reside"),
			config.GetInt("min_mortgage_term"),
			config.GetInt("max_mortgage_term"),
		)(csvc)
	}

	calculatorGrpcServer := calculatorGrpc.MakeHandler(csvc, logger, config.GetDuration("rate_limit_duration"))

	grpcAddr := strings.Join([]string{":", config.GetString("GRPC_PORT")}, "")

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		level.Error(logger).Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterCalculatorServer(baseServer, calculatorGrpcServer)
		level.Info(logger).Log("msg", "Server started successfully")
		level.Info(logger).Log("transport", "gRPC", "addr", grpcAddr)
		errs <- baseServer.Serve(grpcListener)
	}()

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	level.Error(logger).Log("exit", <-errs)
}
