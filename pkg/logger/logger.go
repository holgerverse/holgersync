package logger

import (
	"log"
	"os"

	"github.com/holgerverse/holgersync/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

type CmdLogger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

func NewCliLogger(cfg *config.Config) *CmdLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &CmdLogger{cfg: cfg}
}

// LoggerLevelMap maps the log level to the zapcore.Level
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// Return the zapcore log level specified by the string
func (l *CmdLogger) getLoggerLevel(logLevel string) zapcore.Level {

	level, exist := loggerLevelMap[logLevel]
	if !exist {
		return zapcore.DebugLevel
	}

	return level

}

func (l *CmdLogger) InitLogger() {

	// Check wether the log level is valid and then set it, if not correct set debug as default
	logLevel := l.getLoggerLevel(l.cfg.Logger.Level)

	// Create a new encoder config for the logger and define default values
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.CallerKey = "caller"

	// Create encoders
	consoleEncoder := zapcore.NewConsoleEncoder(encoder)
	fileEncoder := zapcore.NewJSONEncoder(encoder)

	// Create the log file
	logFile, err := os.Create(l.cfg.Logger.Destination)
	if err != nil {
		log.Fatal("Could not create log file")
	}

	// Create logging sinks
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(logFile), logLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(os.Stderr), logLevel),
	)

	// Apply cores to the logger
	logger := zap.New(core)

	// Pass configuration to the logger
	l.sugarLogger = logger.Sugar()

}

func (l *CmdLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *CmdLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *CmdLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *CmdLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *CmdLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *CmdLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *CmdLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *CmdLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *CmdLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *CmdLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *CmdLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *CmdLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *CmdLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *CmdLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
