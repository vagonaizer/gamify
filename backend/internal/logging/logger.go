package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init(level string) error {
	logLevel := parseLevel(level)

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "",
		LevelKey:       "level",
		NameKey:        "",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customLevelEncoder,
		EncodeCaller:   customCallerEncoder,
		EncodeTime:     nil,
		EncodeDuration: nil,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel)

	Log = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	zap.ReplaceGlobals(Log)
	return nil
}

func parseLevel(level string) zapcore.Level {
	lvl, err := zapcore.ParseLevel(level)
	if err != nil {
		return zapcore.InfoLevel
	}
	return lvl
}

// Кастомный энкодер для уровней
func customLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var emoji string
	switch l {
	case zapcore.DebugLevel:
		emoji = "🐛 DEBUG"
	case zapcore.InfoLevel:
		emoji = "ℹ️ INFO "
	case zapcore.WarnLevel:
		emoji = "⚠️ WARN "
	case zapcore.ErrorLevel:
		emoji = "❌ ERROR"
	case zapcore.FatalLevel:
		emoji = "💀 FATAL"
	default:
		emoji = "🔷 UNKWN"
	}
	enc.AppendString(fmt.Sprintf("%-2s", emoji)) // Меньше расстояния, чтобы выровнять
}

// Кастомный энкодер для caller
func customCallerEncoder(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("📍 %-2s", c.TrimmedPath())) // Выравнивание пути
}
