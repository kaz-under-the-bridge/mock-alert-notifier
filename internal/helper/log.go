package helper

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// declare logger as global variable
var Logger *slog.Logger

// get logger with log/slog
func GetNewLogger(ctx context.Context) error {
	var err error
	var logType, logFile, logFormat string
	// get config from ctx by key logType string
	// logType are "file" or "stdout"
	// logFormat are "json" or "text"
	logType = GetLogType(ctx)
	if logType == "test" {
		logFile = "stdout"
		logFormat = "text"
	} else {
		logFile = GetLogFile(ctx)
		logFormat = GetLogFormat(ctx)
	}

	switch logType {
	case "test":
		Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	case "file":
		if logFile == "" {
			return fmt.Errorf("logFile is empty")
		}

		// create logger with file output
		var file *os.File
		if file, err = os.Open(logFile); nil != err {
			return err
		}

		// if logFormat is json, create logger with json format
		if logFormat == "json" {
			Logger = slog.New(slog.NewJSONHandler(file, nil))
		} else {
			// if logFormat is text, create logger with text format
			Logger = slog.New(slog.NewTextHandler(file, nil))
		}

	case "stdout":
		// create logger with stdout output
		// if logFormat is json, create logger with json format
		if logFormat == "json" {
			Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		} else {
			// if logFormat is text, create logger with text format
			Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
		}

	default:
		return fmt.Errorf("logType(%s) is not supported", logType)
	}

	return nil
}
