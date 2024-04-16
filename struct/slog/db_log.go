package slog

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type log struct {
	logger.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func (l *log) LogMode(level logger.LogLevel) logger.Interface {
	Keylogger := *l
	Keylogger.LogLevel = level
	return &Keylogger
}
func (l *log) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
	}
}
func (l *log) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
	}
}
func (l *log) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
	}
}

func (l *log) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > logger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && !errors.Is(err, gorm.ErrRecordNotFound) && l.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				fmt.Println(fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				fmt.Println(fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				fmt.Println(fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				fmt.Println(fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		case l.LogLevel == logger.Info:
			sql, rows := fc()
			if rows == -1 {
				fmt.Println(fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				fmt.Println(fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		}
	}
}

func NewLogger() *log {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	return &log{
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}
