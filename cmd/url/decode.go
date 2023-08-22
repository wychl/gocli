/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application foo.
*/
package urlcmd

import (
	"io"
	"log"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "url decode",
	Long:  `url decode`,
	Run: func(cmd *cobra.Command, args []string) {
		var data []byte
		if len(args) > 0 {
			data = []byte(args[0])
		} else {
			var err error
			data, err = io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(data) > 0 {
			str, _ := url.QueryUnescape(string(data))
			os.Stdout.Write([]byte(str))
		}
	},
}

func init() {
	URLCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
