package logger

import "go.uber.org/zap"

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type ZapLogger struct {
	sugarLogger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &ZapLogger{
		sugarLogger: logger.Sugar(),
	}
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *ZapLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *ZapLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *ZapLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *ZapLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
