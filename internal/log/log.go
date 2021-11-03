package log

import (
	"github.com/mabta/clpc/internal/cfg"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//var Logger *zap.Logger
var Logger *zap.SugaredLogger

func Init(config *cfg.LogConfig) error {
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(config.Level)); err != nil {
		return err
	}
	core := zapcore.NewCore(getEncoder(), getLogWriter(config), l)
	Logger = zap.New(core, zap.AddCaller()).Sugar()
	return nil
}

func getLogWriter(config *cfg.LogConfig) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	return zapcore.AddSync(lumberjackLogger)
}
func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	//return zapcore.NewJSONEncoder(cfg)
	return zapcore.NewConsoleEncoder(cfg)
}
