package app

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func InitLogger(Debug bool, serviceName string) log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		if Debug {
			logger = level.NewFilter(logger, level.AllowDebug())
		}
		logger = log.With(logger,
			"svc", serviceName,
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	return logger
}
