package helper

import "context"

type contextKey string

const (
	appNameKey   contextKey = "appName"
	logTypeKey   contextKey = "logType"
	logFileKey   contextKey = "logFile"
	logFormatKey contextKey = "logFormat"
)

// setter and gettter for context value

// for appName
func SetAppName(ctx context.Context, appName string) context.Context {
	return context.WithValue(ctx, appNameKey, appName)
}

func GetAppName(ctx context.Context) string {
	return ctx.Value(appNameKey).(string)
}

// for LogType
func SetLogType(ctx context.Context, logType string) context.Context {
	return context.WithValue(ctx, logTypeKey, logType)
}

func GetLogType(ctx context.Context) string {
	return ctx.Value(logTypeKey).(string)
}

// for LogFile
func SetLogFile(ctx context.Context, logFile string) context.Context {
	return context.WithValue(ctx, logFileKey, logFile)
}

func GetLogFile(ctx context.Context) string {
	return ctx.Value(logFileKey).(string)
}

// for LogFormat
func SetLogFormat(ctx context.Context, logFormat string) context.Context {
	return context.WithValue(ctx, logFormatKey, logFormat)
}

func GetLogFormat(ctx context.Context) string {
	return ctx.Value(logFormatKey).(string)
}
