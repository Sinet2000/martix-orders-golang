package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log     *zap.Logger
	once    sync.Once
	logConf *LoggerConfig
)

// InitLogger initializes the logger with configuration from the JSON file
func InitLogger(environment string) error {
	var err error
	once.Do(func() {
		err = setupLogger(environment)
	})
	return err
}

func setupLogger(environment string) error {
	// Load configuration
	if err := loadConfig(); err != nil {
		return fmt.Errorf("failed to load logger config: %w", err)
	}

	// Choose configuration based on environment
	var cfg LoggerEnvironmentConfig
	if environment == "production" {
		cfg = logConf.Production
	} else {
		cfg = logConf.Development
	}

	// Create basic encoder config
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     cfg.EncoderConfig.MessageKey,
		LevelKey:       cfg.EncoderConfig.LevelKey,
		TimeKey:        cfg.EncoderConfig.TimeKey,
		NameKey:        "logger",
		CallerKey:      cfg.EncoderConfig.CallerKey,
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		EncodeLevel:    getEncoderByName(cfg.EncoderConfig.LevelEncoder),
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// Create Zap configuration
	zapConfig := zap.Config{
		Level:            getLogLevel(cfg.Level),
		Development:      environment != "production",
		Encoding:         cfg.Encoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      cfg.OutputPaths,
		ErrorOutputPaths: cfg.ErrorOutputPaths,
	}

	// Build the logger
	logger, err := zapConfig.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return fmt.Errorf("failed to build logger: %w", err)
	}

	log = logger
	return nil
}

func loadConfig() error {
	data, err := os.ReadFile("configs/logger_config.json")
	if err != nil {
		return fmt.Errorf("failed to read logger config file: %w", err)
	}

	logConf = &LoggerConfig{}
	if err := json.Unmarshal(data, logConf); err != nil {
		return fmt.Errorf("failed to unmarshal logger config: %w", err)
	}

	return nil
}

func getEncoderByName(encoder string) zapcore.LevelEncoder {
	switch encoder {
	case "capitalColor":
		return zapcore.CapitalColorLevelEncoder
	case "capital":
		return zapcore.CapitalLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func getLogLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}

// Logging methods
func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}

// WithFields adds fields to the logger
func WithFields(fields ...zapcore.Field) *zap.Logger {
	return log.With(fields...)
}

// Sync flushes any buffered log entries
func Sync() error {
	return log.Sync()
}
