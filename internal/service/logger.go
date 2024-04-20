package service

import (
	"fmt"

	"quasar-evm/pkg/logger"
)

func logInfo(msg string, method string) {
	logger.Log().WithField("layer", fmt.Sprintf("Service-%s", method)).Info(msg)
}

func logErr(msg string, method string) {
	logger.Log().WithField("layer", fmt.Sprintf("Service-%s", method)).Error(msg)
}
