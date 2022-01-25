package logs

import (
	"github.com/gofiber/fiber/v2"
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
func Logger(ctx *fiber.Ctx, errorMessage string, logLevel string) {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs.json",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
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
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
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
			logger.Debug(errorMessage, zap.String("ipAddress", ctx.IP()), zap.String("responseHeader", ctx.Response().Header.String()))
			break
		}
	case INFO:
		{
			logger.Info(errorMessage, zap.String("ipAddress", ctx.IP()), zap.String("responseHeader", ctx.Response().Header.String()))
			break
		}

	case ERROR:
		{
			logger.Error(errorMessage, zap.String("ipAddress", ctx.IP()), zap.String("responseHeader", ctx.Response().Header.String()))
			break
		}
	case WARN:
		{
			logger.Warn(errorMessage, zap.String("ipAddress", ctx.IP()), zap.String("responseHeader", ctx.Response().Header.String()))
			break
		}
	case FATAL:
		{
			logger.Fatal(errorMessage, zap.String("ipAddress", ctx.IP()), zap.String("responseHeader", ctx.Response().Header.String()))
			break
		}
	}
}
