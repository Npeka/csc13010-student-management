package initialize

import (
	"github.com/csc13010-student-management/config"
	"github.com/csc13010-student-management/pkg/logger"
)

func NewLogger(cfg config.LoggerConfig) *logger.LoggerZap {
	return logger.NewLogger(cfg)
}
