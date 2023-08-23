/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
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
}
