/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application foo.
*/
package urlcmd

import (
	"github.com/spf13/cobra"
)

const urlExp = `
# url编码
gocli url encode https://example.com/你好

# url解码
gocli url decode https://example.com/%E4%BD%A0%E5%A5%BD
`

// URLCmd represents the url command
var URLCmd = &cobra.Command{
	Use:   "url",
	Short: "url",
	Long:  `url`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// urlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// urlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
