/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"log"
	"os"
	"strconv"

	"codeberg.org/gruf/go-bytesize"
	"github.com/spf13/cobra"
)

var byteSizeF string

const byteSizeExp = `
# 格式化字节大小
gocli bytesize 1024

# 格式化文件内容字节大小
gocli bytesize -f test.txt
`

// bytesizeCmd represents the bytesize command
var bytesizeCmd = &cobra.Command{
	Use:     "bytesize",
	Short:   "格式化字节大小",
	Long:    `格式化字节大小`,
	Example: byteSizeExp,
	Run: func(cmd *cobra.Command, args []string) {
		sizeStr := "0"
		if file := cmd.Flags().Lookup("file").Value.String(); file != "" {
			info, err := os.Stat(file)
			if err != nil {
				log.Fatal(err)
			}
			sizeStr = strconv.Itoa(int(info.Size()))
		} else if len(args) > 0 {
			sizeStr = args[0]
		}
		bSize, err := bytesize.ParseSize(sizeStr)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.WriteString(bSize.String())
	},
}

func init() {
	rootCmd.AddCommand(bytesizeCmd)
	bytesizeCmd.Flags().StringVarP(&byteSizeF, "file", "f", "", "文件路径")
}
