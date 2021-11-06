package main

import (
	"github.com/GoPlugin/Plugin/core/logger"
	"github.com/GoPlugin/external-initiator/integration/mock-client/grpc"
	"github.com/GoPlugin/external-initiator/integration/mock-client/web"
	"go.uber.org/zap/zapcore"
)

func init() {
	logger.SetLogger(logger.CreateProductionLogger("", false, zapcore.DebugLevel, false))
}

func main() {
	logger.Info("Starting mock blockchain client")

	grpc.RunServer()
	web.RunWebserver()
}
