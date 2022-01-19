package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DEBUG string = "DEBUG"
	INFO  string = "INFO"
	FATAL string = "FATAL"
	WARN  string = "WARN"
	ERROR string = "ERROR"
)

func Logger(errorMessage string, logLevel string, ipAddress string) {

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, loggerError := cfg.Build()

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
			logger.Error(errorMessage, zap.String("ipAddress", ipAddress))
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
