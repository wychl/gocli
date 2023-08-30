/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// qrcodeCmd represents the qrcode command
var qrcodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "二维码",
	Long:  `二维码`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(qrcodeCmd)
}
