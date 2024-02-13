package database

// import (
// 	"context"
// 	"strings"
// 	"time"

// 	sharecycle_logger "backend/logger"
// 	"regexp"

// 	"go.uber.org/zap/zapcore"
// 	"gorm.io/gorm/logger"
// )

// type gormLogger struct {
// 	logger.Config
// 	infoStr, warnStr, errStr string
// 	isFormat                 bool
// }

// func NewGormLogger(logLevel logger.LogLevel, isFormat bool) logger.Interface {
// 	config := logger.Config{
// 		SlowThreshold:             time.Second,
// 		IgnoreRecordNotFoundError: true,
// 		ParameterizedQueries:      true,
// 		LogLevel:                  logLevel,
// 	}

// 	return &gormLogger{
// 		Config:   config,
// 		infoStr:  "[info] ",
// 		warnStr:  "[warn] ",
// 		errStr:   "[error] ",
// 		isFormat: isFormat,
// 	}
// }

// type QueryLog struct {
// 	SQL       string
// 	Rows      int64
// 	ElapsedMs float64
// 	Caller    string
// }

// var (
// 	multiSpaceRegexp = regexp.MustCompile(`\s{2,}`)
// )

// func (o *QueryLog) FormatLoggingSQL() {
// 	o.SQL = strings.ReplaceAll(o.SQL, "\n", " ")
// 	o.SQL = strings.ReplaceAll(o.SQL, "\t", " ")
// 	o.SQL = multiSpaceRegexp.ReplaceAllString(o.SQL, " ")
// }

// func (o *QueryLog) MarshalLogObject(enc zapcore.ObjectEncoder) error {
// 	enc.AddString("sql", o.SQL)
// 	enc.AddInt64("rows", o.Rows)
// 	enc.AddFloat64("elapsed_ms", o.ElapsedMs)
// 	enc.AddString("caller", o.Caller)
// 	return nil
// }

// func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
// 	newLogger := *l
// 	newLogger.LogLevel = level
// 	return &newLogger
// }

// func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
// 	ctxLogger := sharecycle_logger.FromContext(ctx)

// 	if l.LogLevel >= logger.Info {
// 		ctxLogger.With("db_log_caller", utils.FileWithLineNum()).Infof(l.infoStr+msg, data...)
// 	}
// }
