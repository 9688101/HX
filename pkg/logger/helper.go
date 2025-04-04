package logger

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/9688101/HX/pkg/helper"
	"go.uber.org/zap"
)

// extractContextField 从 context 中提取请求 ID（如果存在）
func extractContextField(ctx context.Context) zap.Field {
	if ctx != nil {
		if reqID := helper.GetRequestID(ctx); reqID != "" {
			return zap.String("requestId", reqID)
		}
	}
	return zap.Skip()
}

// getCallerInfo 获取调用者文件、行号及函数名称信息
func getCallerInfo(skip int) (string, string) {
	funcName := "[unknown] "
	pc, file, line, ok := runtime.Caller(skip)
	if ok && pc != 0 {
		if fn := runtime.FuncForPC(pc); fn != nil {
			parts := strings.Split(fn.Name(), ".")
			funcName = "[" + parts[len(parts)-1] + "] "
		}
	} else {
		file = "unknown"
		line = 0
	}
	// 可选：简化文件路径显示
	parts := strings.Split(file, "one-api/")
	if len(parts) > 1 {
		file = parts[1]
	}
	return fmt.Sprintf(" | %s:%d", file, line), funcName
}

// ToZapField 根据 value 的具体类型将 key-value 对转换为 zap.Field
func ToZapField(key string, value interface{}) zap.Field {
	switch v := value.(type) {
	case string:
		return zap.String(key, v)
	case int:
		return zap.Int(key, v)
	case int32:
		return zap.Int32(key, v)
	case int64:
		return zap.Int64(key, v)
	case float32:
		return zap.Float32(key, v)
	case float64:
		return zap.Float64(key, v)
	case bool:
		return zap.Bool(key, v)
	case error:
		return zap.Error(v)
	default:
		// 如果类型不在上述范围内，则使用 Any 方法进行处理
		return zap.Any(key, v)
	}
}

// ToZapFields 将 map[string]interface{} 中的所有键值对转换为 zap.Field 切片
func ToZapFields(data map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(data))
	for k, v := range data {
		fields = append(fields, ToZapField(k, v))
	}
	return fields
}
