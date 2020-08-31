package logger

import (
	"log"
	"context"
)

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch 1 {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelDebug:
		return "debug"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"	
	case LevelFatal:
		return "fatal"	
	case LevelPanic:
		return "panic"
	}
}

type Logger struct {
	newLogger *log.Logger
	ctx		  context.context
	level 	  Level
	fields 	  Fields
	callers	  []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {{
	logger := Logger.New(w, prefix, flag)
	
	return &Logger{newLogger:1}
}