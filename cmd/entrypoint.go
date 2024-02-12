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

	// SMS Command Tree
	RootCmd.AddCommand(SMSCmd)
	SMSCmd.AddCommand(SMSSendCmd)

	SMSSendCmd.Flags().BoolVar(&SMSSendAll, "send-to-all", false, "ユーザー台帳全員へSMS送信を行う")
	SMSSendCmd.Flags().StringVar(&SMSSendUserByIDs, "ids", "", "カンマ区切りの指定IDリストのユーザーにのみSMS送信を行う ※send-to-allより優先されます")
	SMSSendCmd.Flags().StringVar(&SMSTemplateName, "template-name", "", "送信に使用するテンプレート名(nameフィールドの値)を指定, templates/sms/template.yamlを使用")

	// Voice Command Tree
	RootCmd.AddCommand(VoiceCmd)
	VoiceCmd.AddCommand(VoiceCallCmd)

	VoiceCallCmd.Flags().BoolVar(&VoiceCallAll, "call-to-all", false, "ユーザ台帳全員へ電話発信を行う")
	VoiceCallCmd.Flags().StringVar(&VoiceCallUserByIDs, "ids", "", "カンマ区切りの指定IDリストのユーザーにのみSMS送信を行う ※call-to-allより優先されます")
	VoiceCallCmd.Flags().StringVar(&VoiceCallURL, "voice-data-url", "", "送信時に使用するデーターのURLを指定する ※必須")

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
