package logs

import (
	"github.com/natefinch/lumberjack"
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

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}
func Logger(errorMessage string, logLevel string, ipAddress string) {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs.json",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	})
	cfg := zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "@timestamp",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		w,
		zap.DebugLevel,
	)
	logger := zap.New(core)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	// ECS Zap Log
	//encoderConfig := ecszap.NewDefaultEncoderConfig()
	//core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	//logger := zap.New(core, zap.AddCaller())

	switch logLevel {
	case DEBUG:
		{
			logger.Debug(errorMessage, zap.Stack("stack"), zap.String("ipAddress", ipAddress))
			break
		}
	case INFO:
		{
			logger.Info(errorMessage, zap.Stack("stack"), zap.String("ipAddress", ipAddress))
			break
		}
	case ERROR:
		{
			logger.Error(errorMessage, zap.Stack("stack"), zap.String("ipAddress", ipAddress))
			break
		}
	case WARN:
		{
			logger.Warn(errorMessage, zap.Stack("stack"), zap.String("ipAddress", ipAddress))
			break
		}
	case FATAL:
		{
			logger.Fatal(errorMessage, zap.Stack("stack"), zap.String("ipAddress", ipAddress))
			break
		}
	}
}
