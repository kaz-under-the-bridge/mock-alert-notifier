package helper

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// declare logger as global variable
var logger *slog.Logger

// get logger with log/slog
func GetNewLogger(ctx context.Context) (*slog.Logger, error) {
	var err error
	// get config from ctx by key logType string
	// logType are "file" or "stdout"
	// logFormat are "json" or "text"
	logType := GetLogType(ctx)
	logFile := GetLogFile(ctx)
	logFormat := GetLogFormat(ctx)

	// return logger if logger is already exist
	if logger != nil {
		return logger, nil
	}

	switch logType {
	case "file":
		if logFile == "" {
			return nil, fmt.Errorf("logFile is empty")
		}

		// create logger with file output
		var file *os.File
		if file, err = os.Open(logFile); nil != err {
			return nil, err
		}

		// if logFormat is json, create logger with json format
		if logFormat == "json" {
			logger = slog.New(slog.NewJSONHandler(file, nil))
		} else {
			// if logFormat is text, create logger with text format
			logger = slog.New(slog.NewTextHandler(file, nil))
		}

	case "stdout":
		// create logger with stdout output
		// if logFormat is json, create logger with json format
		if logFormat == "json" {
			logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		} else {
			// if logFormat is text, create logger with text format
			logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
		}

	default:
		return nil, fmt.Errorf("logType is not supported")
	}

	return logger, nil
}
