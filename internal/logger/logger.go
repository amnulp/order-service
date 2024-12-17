package logger

import "go.uber.org/zap"

const noSuchEnvErrorMsg = "no such environment"

var loggerInstance *zap.SugaredLogger

func InitLogger(env string) {
	if loggerInstance != nil {
		return
	}

	switch env {
	case "PROD":
		loggerInstance = zap.Must(
			zap.NewProduction(),
		).Sugar()
	case "DEV":
		loggerInstance = zap.Must(
			zap.NewDevelopment(),
		).Sugar()
	default:
		panic(noSuchEnvErrorMsg)
	}

}

func Warn(args ...interface{}) {
	loggerInstance.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	loggerInstance.Warnf(template, args...)
}

func Error(args ...interface{}) {
	loggerInstance.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	loggerInstance.Errorf(template, args...)
}

func Info(args ...interface{}) {
	loggerInstance.Info(args...)
}

func Infof(template string, args ...interface{}) {
	loggerInstance.Infof(template, args...)
}

func Debug(args ...interface{}) {
	loggerInstance.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	loggerInstance.Debugf(template, args...)
}

func Fatal(args ...interface{}) {
	loggerInstance.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	loggerInstance.Fatalf(template, args...)
}
