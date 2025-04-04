package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/9688101/HX/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ──────────────
// 基于 zap 的日志实现
// ─────────────────

func newZapLogger(cfg *config.LoggerConfig) (Logger, error) {
	var logPath string
	if cfg.LogDir != "" {
		if cfg.OnlyOneLogFile {
			logPath = filepath.Join(cfg.LogDir, "HX.log")
		} else {
			logPath = filepath.Join(cfg.LogDir, fmt.Sprintf("HX-%s.log", time.Now().Format("20060102")))
		}
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 默认输出到标准输出
	var writer zapcore.WriteSyncer = zapcore.AddSync(os.Stdout)
	if logPath != "" {
		f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		writer = zapcore.AddSync(io.MultiWriter(os.Stdout, f))
		// 同步 gin 默认输出
		gin.DefaultWriter = writer
		gin.DefaultErrorWriter = writer
	}

	level := zapcore.InfoLevel
	if cfg.DebugEnabled {
		level = zapcore.DebugLevel
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		level,
	)
	// 添加调用者信息，跳过1层包装
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &zapLoggerImpl{logger: zapLogger}, nil
}

type zapLoggerImpl struct {
	logger *zap.Logger
}

func (z *zapLoggerImpl) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append([]zap.Field{extractContextField(ctx)}, fields...)
	z.logger.Debug(msg, fields...)
}

func (z *zapLoggerImpl) Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append([]zap.Field{extractContextField(ctx)}, fields...)
	z.logger.Info(msg, fields...)
}

func (z *zapLoggerImpl) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append([]zap.Field{extractContextField(ctx)}, fields...)
	z.logger.Warn(msg, fields...)
}

func (z *zapLoggerImpl) Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append([]zap.Field{extractContextField(ctx)}, fields...)
	z.logger.Error(msg, fields...)
}

func (z *zapLoggerImpl) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append([]zap.Field{extractContextField(ctx)}, fields...)
	z.logger.Fatal(msg, fields...)
}
