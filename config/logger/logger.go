package logger

import (
	"encoding/json"
	"fmt"
	"github.com/Sinet2000/Martix-Orders-Go/config"
	"github.com/Sinet2000/Martix-Orders-Go/pkg/fileutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
)

type LoggerConfig struct {
	Level            string                `json:"level"`
	Encoding         string                `json:"encoding"`
	OutputPaths      []string              `json:"outputPaths"`
	ErrorOutputPaths []string              `json:"errorOutputPaths"`
	EncoderConfig    zapcore.EncoderConfig `json:"encoderConfig"`
}

func ConfigureZapLogger() (*zap.Logger, error) {
	env := config.GetEnv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envConfig, err := loadConfigByEnv(env)
	if err != nil {
		return nil, err
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(parseLogLevel(envConfig.Level)),
		Development:      env == "development",
		Encoding:         envConfig.Encoding,
		OutputPaths:      envConfig.OutputPaths,
		ErrorOutputPaths: envConfig.ErrorOutputPaths,
		EncoderConfig:    envConfig.EncoderConfig,
	}
	return zapConfig.Build()
}

func loadConfigByEnv(env string) (*LoggerConfig, error) {
	configFilePath := filepath.Join(".", "logger_config.json")
	fileReader := &fileutil.FileSystemReader{}
	data, err := fileReader.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read logger config: %v", err)
	}

	loadedConfigs := make(map[string]LoggerConfig)
	err = json.Unmarshal(data, &loadedConfigs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse logger config: %v", err)
	}

	loadedEnvConfig, exists := loadedConfigs[env]
	if !exists {
		return nil, fmt.Errorf("logger config for environment '%s' not found", env)
	}

	return &loadedEnvConfig, nil
}

func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}
