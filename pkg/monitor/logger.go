package monitor

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

var (
	infoLogPath = "../log/blog.info"
	warnLogPath = "../log/blog.warn"
	openDebug   = true
)

//Initlog 初始化log.debug和info输出一个文件，warn，error输出一个文件
func InitLog() {
    config := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }
    infoLog := zapcore.AddSync(&lumberjack.Logger{
        Filename:   infoLogPath,
        MaxSize:    100, // megabytes
        MaxBackups: 10000,
        MaxAge:     7, // days
    })
    warnLog := zapcore.AddSync(&lumberjack.Logger{
        Filename:   warnLogPath,
        MaxSize:    500, // megabytes
        MaxBackups: 10000,
        MaxAge:     7, // days
    })
    core := zapcore.NewTee(
        zapcore.NewCore(
            zapcore.NewJSONEncoder(config),
            infoLog,
            zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
                if openDebug {
                    return lvl <= zapcore.InfoLevel
                }
                return lvl == zapcore.InfoLevel
            },
            ),
        ),
        zapcore.NewCore(
            zapcore.NewJSONEncoder(config),
            warnLog,
            zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
                return lvl >= zapcore.WarnLevel
            },
            ),
        ),
    )

    logger := zap.New(core)
    logger = logger.WithOptions(zap.AddCaller())
    zap.ReplaceGlobals(logger)
}