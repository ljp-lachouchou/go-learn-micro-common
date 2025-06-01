package common

import (
	"go.uber.org/zap"                  // Zap 核心日志库[1,3,5](@ref)
	"go.uber.org/zap/zapcore"          // Zap 核心编码和输出控制[1,5](@ref)
	"gopkg.in/natefinch/lumberjack.v2" // 日志切割归档库[4,7,8](@ref)
)

var (
	logger *zap.SugaredLogger // 全局 SugaredLogger 实例（支持结构化+printf风格）[1,5](@ref)
)

func init() {
	fileName := "micro.log" // 日志文件名[4,7](@ref)

	// 配置 Lumberjack 日志切割
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName, // 日志文件路径
		MaxSize:    512,      // 单文件最大 512MB（超过则切割）[4,7,8](@ref)
		MaxBackups: 1,        // 保留最多 1 个旧日志备份[4,7](@ref)
		LocalTime:  true,     // 使用本地时间命名备份文件[7](@ref)
		Compress:   true,     // 压缩旧日志节省空间[4,7,8](@ref)
	})

	// 编码器配置
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式化为 ISO8601 标准（如 "2025-06-01T12:00:00Z"）[1,5](@ref)

	// 创建 Zap 核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),      // JSON 格式编码器[1,5](@ref)
		syncWriter,                           // 输出目标（Lumberjack）
		zap.NewAtomicLevelAt(zap.DebugLevel), // 日志级别：Debug 及以上[2,5](@ref)
	)

	// 构建 Logger 实例
	log := zap.New(
		core,
		zap.AddCaller(),      // 自动添加调用者信息（文件名+行号）[2,5](@ref)
		zap.AddCallerSkip(1), // 跳过一层调用栈（避免显示本封装函数）[5](@ref)
	)
	logger = log.Sugar() // 转换为 SugaredLogger（简化调用）[1,5](@ref)
}
func Info(args ...interface{}) {
	logger.Info(args...)
}
func Warn(args ...interface{}) {
	logger.Warn(args...)
}
func Error(args ...interface{}) {
	logger.Error(args...)
}
func Panic(args ...interface{}) {
	logger.Panic(args...)

}
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
