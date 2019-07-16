package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var instance *zap.SugaredLogger


func init() {
	//Example和Production使用的是json格式输出，Development使用行的形式输出
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	zap, err := config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	instance = zap.Sugar()
}

func Debug(args ...interface{}) {
	instance.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	instance.Debugf(template, args...)
}

func Info(args ...interface{}) {
	instance.Info(args)
}

func Infof(template string, args ...interface{}) {
	instance.Infof(template, args...)
}

func Warn(args ...interface{}) {
	instance.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	instance.Warnf(template, args...)
}

func Error(args ...interface{}) {
	instance.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	instance.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	instance.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	instance.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	instance.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	instance.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	instance.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	instance.Fatalf(template, args...)
}
