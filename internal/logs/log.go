package logs

import (
	"go.uber.org/zap"
)

const (
	DEBUG string = "DEBUG"
	INFO  string = "INFO"
	FATAL string = "FATAL"
	WARN  string = "WARN"
	ERROR string = "ERROR"
)

func Logger(errorMessage string, logLevel string, err error) {
	logger, loggerError := zap.NewProduction()

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
	if loggerError != nil {
		return
	}
	switch logLevel {
	case DEBUG:
		{
			logger.Debug(errorMessage)
			break
		}
	case INFO:
		{
			logger.Info(errorMessage)
			break
		}
	case ERROR:
		{
			logger.Error(errorMessage, zap.String("err", err.Error()))
			break
		}
	case WARN:
		{
			logger.Warn(errorMessage)
			break
		}
	case FATAL:
		{
			logger.Fatal(errorMessage)
			break
		}
	}
}
