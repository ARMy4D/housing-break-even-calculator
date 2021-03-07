package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log/level"

	"github.com/army4d/housing-break-even-calculator/app"
	"github.com/army4d/housing-break-even-calculator/services/calculator"
	calculatorSvc "github.com/army4d/housing-break-even-calculator/services/calculator/service/client"
)

const (
	defaultConfigName = "default"
	serviceName       = "calculator_client"
)

func main() {

	calcFlags := flag.NewFlagSet("calculate", flag.ExitOnError)

	var (
		rent            = calcFlags.Float64("rent", 0, "specify rent")
		rentIcreaseRate = calcFlags.Float64("rent-inc-rate", 0, "specify rent increase rate")
		downPayment     = calcFlags.Float64("down-payment", 0, "specify down payment")
		intrest         = calcFlags.Float64("intrest", 0, "specify intrest rate")
		term            = calcFlags.Int("term", 0, "specify term")
		price           = calcFlags.Float64("price", 0, "specify price")
		pTax            = calcFlags.Float64("p-tax", 0, "specify propery tax rate")
		tTax            = calcFlags.Float64("t-tax", 0, "specify propery transfer tax rate")
		reside          = calcFlags.Int("res", 0, "specify expected residance in years")
	)

	calcFlags.Parse(os.Args[1:])

	config, err := app.InitConfig(defaultConfigName)
	if err != nil {
		panic(fmt.Errorf("fatal error in config file: %v+", err))
	}
	logger := app.InitLogger(config.GetBool("debug"), serviceName)

	grpcAddr := strings.Join([]string{config.GetString("GRPC_HOST"), ":", config.GetString("GRPC_PORT")}, "")
	level.Info(logger).Log("transport", "gRPC", "addr", grpcAddr)

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	var csvc calculator.Service
	{
		csvc = calculatorSvc.NewService(
			logger,
			conn,
			config.GetDuration("rate_limit_duration"),
			app.ERR_NL,
		)
	}

	ctx := context.Background()

	if resp, err := csvc.Calculate(ctx, calculator.RentSetting{
		Rent:                   float32(*rent),
		YearlyRentIncreaseRate: float32(*rentIcreaseRate),
	}, calculator.MortgageSetting{
		DownPayment: float32(*downPayment),
		IntrestRate: float32(*intrest),
		Term:        *term,
	}, calculator.HouseSetting{
		Price:                   float32(*price),
		PropertyTaxRate:         float32(*pTax),
		PropertyTransferTaxRate: float32(*tTax),
		YearsExpectedToReside:   *reside,
	}); err == nil {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Type", "Cost Over Residance Period", "Break Even", "Recommended"})
		t.AppendRow(table.Row{
			"Rent", resp.CostOfRentOverResidancePeriod,
		})
		t.AppendSeparator()
		t.AppendRows([]table.Row{
			{"Mortgage Intrest Only", resp.CostOfMortgageIntrestOverResidancePeriod, resp.BreakEvenYearIntrest, resp.BestPaymentTypeIntrest},
			{"Mortgage Overall", resp.CostOfMortgageOverResidancePeriod, resp.BreakEvenYearOverall, resp.BestPaymentTypeOverall},
		})
		t.Render()
	}
}
