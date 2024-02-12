package cmd

import (
	"os"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/spf13/cobra"
)

var SMSCmd = &cobra.Command{
	Use:   "sms",
	Short: "SMS関連のサブコマンド",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var SMSSendAll bool
var SMSSendUserByIDs string
var SMSTemplateName string
var SMSSendCmd = &cobra.Command{
	Use:   "send",
	Short: "SMS送信",
	Run: func(cmd *cobra.Command, args []string) {
		// テンプレートの指定は必須
		if SMSTemplateName == "" {
			_ = cmd.Help()
			os.Exit(1)
		}

		if SMSSendUserByIDs != "" {
			// ID(カンマ区切り)を対象にSMSを送信
			// err := usecase.SendSMSTo(SMSTemplateName, SMSSendUserByIDs)
			helper.Logger.Info("Send SMS to users by IDs ends successfully")
			os.Exit(0)
		}

		if SMSSendAll {
			// 全ユーザにSMSを送信
			// err := usecase.SendSMSToAllUsers(SMSTemplateName)
			helper.Logger.Info("Send SMS to all users ends successfully")
			os.Exit(0)
		}

	},
}
