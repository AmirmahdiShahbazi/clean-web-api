package logger

import "github.com/AmirmahdiShahbazi/clean-web-api/config"

type Logger interface {
	Init()

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger() Logger{
	loggerCfg := config.GetConfigs()
	if loggerCfg.Logger.Logger == "zap"{
		return newZapLogger(&loggerCfg)
	} else if loggerCfg.Logger.Logger == "zero"{

	}

	panic("logger not supported")
}