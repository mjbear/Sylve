// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

// internal/logger/hclog_adapter.go
package logger

import (
	"io"
	stdlog "log"

	"github.com/hashicorp/go-hclog"
	"github.com/rs/zerolog"
)

// ---------- io.Writer to zerolog ----------
type zerologWriter struct {
	l     zerolog.Logger
	level zerolog.Level
}

func (w zerologWriter) Write(p []byte) (int, error) {
	msg := string(p)
	switch w.level {
	case zerolog.DebugLevel:
		w.l.Debug().Msg(msg)
	case zerolog.WarnLevel:
		w.l.Warn().Msg(msg)
	case zerolog.ErrorLevel:
		w.l.Error().Msg(msg)
	default:
		w.l.Info().Msg(msg)
	}
	return len(p), nil
}

// StandardWriterAdapter is handy for Raft transport/snapshot writers.
func StandardWriterAdapter(zl zerolog.Logger) io.Writer {
	return zerologWriter{l: zl, level: zerolog.InfoLevel}
}

// ---------- hclog.Logger adapter ----------
type ZerologHCLog struct {
	zl          zerolog.Logger
	name        string
	level       hclog.Level
	impliedArgs []interface{}
}

func NewZerologHCLog(zl zerolog.Logger, name string) hclog.Logger {
	return &ZerologHCLog{
		zl:    zl,
		name:  name,
		level: hclog.Info,
	}
}

// Core logging
func (l *ZerologHCLog) Log(level hclog.Level, msg string, args ...interface{}) {
	if !l.accept(level) {
		return
	}
	ev := l.baseEvent(level)
	fields := kvsToMap(append(l.impliedArgs, args...)...)
	ev.Fields(fields).Msg(msg) // <-- pass a single map
}

func (l *ZerologHCLog) Trace(msg string, args ...interface{}) { l.Log(hclog.Trace, msg, args...) }
func (l *ZerologHCLog) Debug(msg string, args ...interface{}) { l.Log(hclog.Debug, msg, args...) }
func (l *ZerologHCLog) Info(msg string, args ...interface{})  { l.Log(hclog.Info, msg, args...) }
func (l *ZerologHCLog) Warn(msg string, args ...interface{})  { l.Log(hclog.Warn, msg, args...) }
func (l *ZerologHCLog) Error(msg string, args ...interface{}) { l.Log(hclog.Error, msg, args...) }

// Level checks (include IsError for your hclog version)
func (l *ZerologHCLog) IsTrace() bool { return l.level <= hclog.Trace }
func (l *ZerologHCLog) IsDebug() bool { return l.level <= hclog.Debug }
func (l *ZerologHCLog) IsInfo() bool  { return l.level <= hclog.Info }
func (l *ZerologHCLog) IsWarn() bool  { return l.level <= hclog.Warn }
func (l *ZerologHCLog) IsError() bool { return l.level <= hclog.Error }

// Naming & scoping
func (l *ZerologHCLog) With(args ...interface{}) hclog.Logger {
	n := *l
	// Add fields to the underlying zerolog logger (single map, no variadic)
	n.zl = l.zl.With().Fields(kvsToMap(args...)).Logger()
	// Track implied args for future calls (as hclog expects)
	n.impliedArgs = append(append([]interface{}{}, l.impliedArgs...), args...)
	return &n
}

func (l *ZerologHCLog) Named(name string) hclog.Logger {
	n := *l
	if l.name != "" {
		n.name = l.name + "." + name
	} else {
		n.name = name
	}
	return &n
}

func (l *ZerologHCLog) ResetNamed(name string) hclog.Logger {
	n := *l
	n.name = name
	return &n
}

func (l *ZerologHCLog) Name() string { return l.name }

// Levels
func (l *ZerologHCLog) SetLevel(level hclog.Level) { l.level = level }
func (l *ZerologHCLog) GetLevel() hclog.Level      { return l.level }

// Standard log/Writer
func (l *ZerologHCLog) StandardLogger(opts *hclog.StandardLoggerOptions) *stdlog.Logger {
	return stdlog.New(l.StandardWriter(opts), "", 0)
}

func (l *ZerologHCLog) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	lev := zerolog.InfoLevel
	if opts != nil {
		switch opts.ForceLevel {
		case hclog.Trace, hclog.Debug:
			lev = zerolog.DebugLevel
		case hclog.Warn:
			lev = zerolog.WarnLevel
		case hclog.Error:
			lev = zerolog.ErrorLevel
		}
	}
	return zerologWriter{l: l.zl.With().Str("subsystem", l.name).Logger(), level: lev}
}

// Implied args
func (l *ZerologHCLog) ImpliedArgs() []interface{} {
	return append([]interface{}{}, l.impliedArgs...)
}

// helpers
func (l *ZerologHCLog) accept(level hclog.Level) bool { return level >= l.level }

func (l *ZerologHCLog) baseEvent(level hclog.Level) *zerolog.Event {
	logger := l.zl.With().Str("subsystem", l.name).Logger()
	switch level {
	case hclog.Trace, hclog.Debug:
		return logger.Debug()
	case hclog.Info:
		return logger.Info()
	case hclog.Warn:
		return logger.Warn()
	case hclog.Error:
		return logger.Error()
	default:
		return logger.Info()
	}
}

func kvsToMap(kvs ...interface{}) map[string]interface{} {
	m := make(map[string]interface{}, len(kvs)/2)
	for i := 0; i < len(kvs); i += 2 {
		// If odd number of args, stash the last value under "arg"
		if i+1 >= len(kvs) {
			m["arg"] = kvs[i]
			break
		}
		key, ok := kvs[i].(string)
		if !ok || key == "" {
			key = "arg"
			// ensure uniqueness if multiple non-string keys
			for {
				if _, exists := m[key]; !exists {
					break
				}
				key += "_"
			}
		}
		m[key] = kvs[i+1]
	}
	return m
}
