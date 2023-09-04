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

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "url encode",
	Long:  `url encode`,
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
			u, err := url.Parse(string(data))
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write([]byte(u.String()))
		}
	},
}

func init() {
	urlCMD.AddCommand(encodeCmd)
}
