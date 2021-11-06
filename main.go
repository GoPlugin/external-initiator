package main

import (
	"github.com/GoPlugin/Plugin/core/logger"
	"github.com/GoPlugin/external-initiator/client"
	"go.uber.org/zap/zapcore"
)

func init() {
	logger.SetLogger(logger.CreateProductionLogger("", false, zapcore.DebugLevel, false))
}

func main() {
	client.Run()
}
