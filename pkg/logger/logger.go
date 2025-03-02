package logger

import (
	"net/http"
	"os"

	"github.com/csc13010-student-management/config"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLoggerTest() *LoggerZap {
	core := zapcore.NewCore(
		GetEncoderLog(),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	return &LoggerZap{zap.New(core, zap.AddCaller())}
}

func NewLogger(cfg config.LoggerConfig) *LoggerZap {
	// debug -> info -> warn -> error -> fatal -> panic
	logLevel := cfg.Log_level
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := GetEncoderLog()

	hook := lumberjack.Logger{
		Filename:   cfg.File_log_name,
		MaxSize:    cfg.Max_size,
		MaxBackups: cfg.Max_backups,
		MaxAge:     cfg.Max_age,
		Compress:   cfg.Compress,
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller())}
}

// format log
func GetEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func GetWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}

func GetRequestID(c *gin.Context) string {
	return c.Request.Header.Get("X-Request-ID")
}

func GetIPAddress(c *gin.Context) string {
	return c.ClientIP()
}

func ErrResponseWithLog(ctx *gin.Context, logger *LoggerZap, err error) {
	logger.Error(
		"ErrResponseWithLog",
		zap.String("RequestID", GetRequestID(ctx)),
		zap.String("IPAddress", GetIPAddress(ctx)),
		zap.Error(err),
	)
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func LogResponseError(ctx *gin.Context, logger *LoggerZap, err error) {
	logger.Error(
		"LogResponseError",
		zap.String("RequestID", GetRequestID(ctx)),
		zap.String("IPAddress", GetIPAddress(ctx)),
		zap.Error(err),
	)
}
