package logger

import (
	"context"
	"fmt"
	"sync"

	"github.com/9688101/HX/config"
	"go.uber.org/zap"
)

// Logger 定义了日志接口，调用方只需依赖此接口
type Logger interface {
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Fatal(ctx context.Context, msg string, fields ...zap.Field)
}

var (
	loggerInstance Logger
	once           sync.Once
)

// InitLogger 初始化全局日志实例，配置从外部传入
func InitLogger(cfg *config.LoggerConfig) error {
	fmt.Println("InitLogger", cfg)
	var err error
	once.Do(func() {
		loggerInstance, err = newZapLogger(cfg)
	})
	return err
}

// GetLogger 返回全局日志实例
func GetLogger() Logger {
	return loggerInstance
}

// ─────────────────────────────
// 包级别便捷函数：直接调用 logger.Info/Debug/Warn/Error/Fatal
// ─────────────────────────────

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	loggerInstance.Debug(ctx, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	loggerInstance.Info(ctx, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	loggerInstance.Warn(ctx, msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	loggerInstance.Error(ctx, msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	loggerInstance.Fatal(ctx, msg, fields...)
}

func SysLog(msg string, fields ...zap.Field) {
	loggerInstance.Info(context.Background(), msg, fields...)
}

func SysDebug(msg string, fields ...zap.Field) {
	loggerInstance.Debug(context.Background(), msg, fields...)
}

func SysWarn(msg string, fields ...zap.Field) {
	loggerInstance.Warn(context.Background(), msg, fields...)
}

func SysError(msg string, fields ...zap.Field) {
	loggerInstance.Error(context.Background(), msg, fields...)
}

func SysFatal(msg string, fields ...zap.Field) {
	loggerInstance.Fatal(context.Background(), msg, fields...)
}
func SysInfo(msg string, fields ...zap.Field) {
	loggerInstance.Info(context.Background(), msg, fields...)
}
