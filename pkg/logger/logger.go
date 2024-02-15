package logger

import (
	"github.com/sirupsen/logrus"
	"sharecycle/configs"

	"go.uber.org/zap"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

var _ Logger = &cLogger{}

type cLogger struct {
	cfg         *configs.Config
	sugarLogger *zap.SugaredLogger
}

func NewArLogger(cfg *configs.Config) *cLogger {

	z, err := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    nil,
	}.Build(zap.AddCaller(), zap.AddCallerSkip(1))

	if err != nil {
		logrus.Fatalln(err)
	}
	z = z.With(
		zap.String("service", cfg.ServiceName),
		zap.String("env", cfg.Env),
	)

	zs := z.Sugar()
	zs.Infof("Init sugar zap foundation")

	return &cLogger{sugarLogger: zs}
}

func (l *cLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *cLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *cLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *cLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *cLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *cLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *cLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *cLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *cLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *cLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
