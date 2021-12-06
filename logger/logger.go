package common

import "go.uber.org/zap"

func InitLogger(debug bool) *zap.Logger {
	var logger *zap.Logger
	var err error
	if debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic("Logger could not initialized")
	}
	zap.ReplaceGlobals(logger)
	return logger
}
