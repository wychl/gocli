/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// pasteCmd represents the past command
var pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "粘贴",
	Long:  `粘贴`,
	Run: func(cmd *cobra.Command, args []string) {
		value, err := clipboard.ReadAll()
		if err != nil {
			log.Fatalln(err)
		}
		_, err = os.Stdout.Write([]byte(value))
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pasteCmd)
}
