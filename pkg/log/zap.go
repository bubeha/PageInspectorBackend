package log

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func NewZapLogger() (*ZapLogger, error) {
	zapCng := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	logger, err := zapCng.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger: logger,
		sugar:  logger.Sugar(),
	}, nil
}

func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

func (l *ZapLogger) Debug(v ...any) {
	l.sugar.Debug(v...)
}

func (l *ZapLogger) Debugf(format string, v ...any) {
	l.sugar.Debugf(format, v...)
}

func (l *ZapLogger) Info(v ...any) {
	l.sugar.Info(v...)
}

func (l *ZapLogger) Infof(format string, v ...any) {
	l.sugar.Infof(format, v...)
}

func (l *ZapLogger) Warn(v ...any) {
	l.sugar.Warn(v...)
}

func (l *ZapLogger) Warnf(format string, v ...any) {
	l.sugar.Warnf(format, v...)
}

func (l *ZapLogger) Error(v ...any) {
	l.sugar.Error(v...)
}

func (l *ZapLogger) Errorf(format string, v ...any) {
	l.sugar.Errorf(format, v...)
}
