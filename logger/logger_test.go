package logger

import (
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name          string
		level         string
		expectedLevel string
	}{
		{"Default Info", "", "INFO"},
		{"Debug Level", "debug", "DEBUG"},
		{"Info Level", "info", "INFO"},
		{"Warn Level", "warn", "WARNING"},
		{"Error Level", "error", "ERROR"},
		{"Critical Level", "critical", "CRITICAL"},
		{"Unknown Level", "unknown", "INFO"}, // Assuming default is INFO
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewLogger(tt.level)
			if level := logger.Level(); strings.ToUpper(level) != tt.expectedLevel {
				t.Errorf("NewLogger(%s) got level %s, want %s", tt.level, level, tt.expectedLevel)
			}
		})
	}
}

func TestLogLevelToStringLevel(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{10, "DEBUG"},
		{20, "INFO"},
		{30, "WARNING"},
		{40, "ERROR"},
		{50, "CRITICAL"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := LogLevelToStringLevel(tt.input)
			if result != tt.expected {
				t.Errorf("LogLevelToStringLevel(%d) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringLevelToLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Debug Level", "debug", 10},
		{"Info Level", "info", 20},
		{"Warn Level", "warn", 30},
		{"Warning Level", "warning", 30},
		{"Error Level", "error", 40},
		{"Critical Level", "critical", 50},
		{"Unknown Level", "unknown", 20}, // Assuming default is INFO (20)
		{"Case Insensitivity", "DeBuG", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringLevelToLogLevel(tt.input)
			if result != tt.expected {
				t.Errorf("StringLevelToLogLevel(%s) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
