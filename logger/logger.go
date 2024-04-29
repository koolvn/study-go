package logger

import (
	"fmt"
	"log"
	"strings"
)

type Logger struct {
	logLevel int
}

func NewLogger(level string) *Logger {
	logLevel := StringLevelToLogLevel(level)
	return &Logger{logLevel: logLevel}
}

func (l *Logger) Level() string {
	return LogLevelToStringLevel(l.logLevel)
}

func (l *Logger) SetLevel(level string) {
	l.logLevel = StringLevelToLogLevel(level)
}

func (l *Logger) Debug(message string) {
	if l.logLevel >= 10 {
		msg := fmt.Sprintf("[DEBUG] %s", message)
		log.Println(msg)
	}
}

func (l *Logger) Info(message string) {
	if l.logLevel >= 20 {
		msg := fmt.Sprintf("[INFO] %s", message)
		log.Println(msg)
	}
}

func (l *Logger) Warning(message string) {
	if l.logLevel >= 30 {
		msg := fmt.Sprintf("[WARNING] %s", message)
		log.Println(msg)
	}
}

func (l *Logger) Error(message string) {
	if l.logLevel >= 40 {
		msg := fmt.Sprintf("[ERROR] %s", message)
		log.Println(msg)
	}
}

func (l *Logger) Critical(message string) {
	if l.logLevel >= 50 {
		msg := fmt.Sprintf("[CRITICAL] %s", message)
		log.Println(msg)
	}
}

func StringLevelToLogLevel(level string) int {
	var logLevel int
	switch strings.ToLower(level) {
	case "debug":
		logLevel = 10
	case "info":
		logLevel = 20
	case "warn":
		logLevel = 30
	case "warning":
		logLevel = 30
	case "error":
		logLevel = 40
	case "critical":
		logLevel = 50
	default:
		logLevel = 20
	}
	return logLevel
}

func LogLevelToStringLevel(level int) string {
	switch {
	case level <= 10:
		return "debug"
	case level <= 20:
		return "info"
	case level <= 30:
		return "warning"
	case level <= 40:
		return "error"
	case level <= 50:
		return "critical"
	default:
		return "unknown"
	}
}
