package cmd

import (
	"os"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/spf13/cobra"
)

var VoiceCmd = &cobra.Command{
	Use:   "voice",
	Short: "電話発信関連のサブコマンド",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var VoiceCallAll bool
var VoiceCallUserByIDs string
var VoiceCallURL string
var VoiceCallCmd = &cobra.Command{
	Use:   "call",
	Short: "電話発信",
	Run: func(cmd *cobra.Command, args []string) {
		if VoiceCallURL == "" {
			_ = cmd.Help()
			os.Exit(1)
		}

		if VoiceCallUserByIDs != "" {
			// ID(カンマ区切り)を対象に電話発信を行う
			// err := usecase.CallTo(VoiceCallURL, VoiceCallUserByIDs)
			helper.Logger.Info("Call to users by IDs ends successfully")
			os.Exit(0)
		}

		if VoiceCallAll {
			// 全ユーザに電話発信を行う
			// err := usecase.CallToAllUsers(VoiceCallURL)
			helper.Logger.Info("Call to users by IDs ends successfully")
			os.Exit(0)
		}

	},
}
