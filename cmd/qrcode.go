/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// qrcodeCmd represents the qrcode command
var qrcodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "二维码",
	Long:  `二维码`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("qrcode called")
	},
}

func init() {
	rootCmd.AddCommand(qrcodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qrcodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qrcodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
