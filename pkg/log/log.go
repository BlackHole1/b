/*
 * SPDX-FileCopyrightText: 2024 Kevin Cui <bh@bugs.cc>
 * SPDX-License-Identifier: MPL-2.0
 */

package log

import (
	"os"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger
var logLevel zap.AtomicLevel

func init() {
	ec := zap.NewDevelopmentEncoderConfig()
	ec.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logLevel = zap.NewAtomicLevel()

	if debugEnv, ok := os.LookupEnv("DEBUG"); ok && debugEnv != "false" {
		logLevel.SetLevel(zap.DebugLevel)
	}

	Log = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(ec),
		zapcore.AddSync(colorable.NewColorableStdout()),
		logLevel,
	))
}

func SetLevel(level zapcore.Level) {
	logLevel.SetLevel(level)
}

func Sync() {
	_ = Log.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
	Sync()
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
	Sync()
}
