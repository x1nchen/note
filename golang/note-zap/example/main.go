package main

import (
	logx "note-zap"
)

func main() {
	logger := logx.NewLogger("example.log", logx.DebugLevel)
	defer logger.Sync()

	logger.Info("info message")
}
