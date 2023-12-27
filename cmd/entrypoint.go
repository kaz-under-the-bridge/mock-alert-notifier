package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/spf13/cobra"
)

var ctx context.Context

func init() {
	ctx = context.Background()

	// get logType from global flag
	RootCmd.PersistentFlags().StringVar(&globalVarLogType, "log-type", "stdout", "stdout or file")
	// get logFile from global flag
	RootCmd.PersistentFlags().StringVar(&globalVarLogFile, "log-file", "", "path to log file")
	// get logFormat from global flag
	RootCmd.PersistentFlags().StringVar(&globalVarLogFormat, "log-format", "text", "text or json")

	setBasicConfigTo(ctx)
}

var Version = "1.0"
var Revision = "00"

var globalVarLogType string
var globalVarLogFile string
var globalVarLogFormat string

var rootVersion bool
var RootCmd = &cobra.Command{
	Use:   "mock-alert-notifier",
	Short: "mock alert notifier",
	Run: func(cmd *cobra.Command, args []string) {
		if rootVersion {
			fmt.Printf("version: %s-%s\n", Version, Revision)
			os.Exit(0)
		}
		_ = cmd.Help()
	},
}

// function to set basic configration to ctx for appName, logType, logFile, logFormat
func setBasicConfigTo(ctx context.Context) context.Context {
	// get logType from global flag
	logType := globalVarLogType
	// get logFile from global flag
	logFile := globalVarLogFile
	// get logFormat from global flag
	logFormat := globalVarLogFormat

	// set basic configration to ctx
	helper.SetAppName(ctx, "mock-alert-notifier")
	helper.SetLogType(ctx, logType)
	if logType == "file" {
		if logFile == "" {
			fmt.Println("log-file is empty, please set log-file path")
			os.Exit(1)
		}
		helper.SetLogFile(ctx, logFile)
	}
	helper.SetLogFormat(ctx, logFormat)

	return ctx
}
