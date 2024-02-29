package logger

import (
	"fmt"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const SERVICE_NAME_FIELD = "service-name"

type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel

	// InvalidLevel is an invalid value for Level.
	//
	// Core implementations may panic if they see messages of this level.
	InvalidLevel = _maxLevel + 1
)

func NewLogger(serviceName string, level Level, environment string) (*zap.Logger, error) {
	logger, err := newConfig(level, environment).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger from default production config: %w", err)
	}
	logger = logger.With(zap.String(SERVICE_NAME_FIELD, serviceName))

	return logger, nil
}

func DefaultLogLevel(environment string) Level {
	if environment == "development" {
		return DebugLevel
	}
	return InfoLevel
}

func newConfig(level Level, environment string) zap.Config {
	isDevEnvironment := environment == "development"
	encoding := "json"
	if isDevEnvironment {
		encoding = "console"
	}
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(level)),
		Development:      isDevEnvironment,
		Encoding:         encoding,
		EncoderConfig:    generateEnvironmentConfig(environment),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		Sampling:         nil,
	}
}

func generateEnvironmentConfig(environment string) zapcore.EncoderConfig {
	cfg := zapdriver.NewDevelopmentEncoderConfig()
	switch environment {
	case "development":
		cfg = zapdriver.NewDevelopmentEncoderConfig()
	case "production":
		cfg = zapdriver.NewProductionEncoderConfig()
	case "staging":
		cfg = zapdriver.NewDevelopmentEncoderConfig()
	case "test":
		cfg = zapdriver.NewDevelopmentEncoderConfig()
	default:
		cfg = zapdriver.NewDevelopmentEncoderConfig()
	}
	cfg.FunctionKey = "func"
	return cfg
}
