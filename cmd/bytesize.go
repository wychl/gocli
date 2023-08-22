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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bytesizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bytesizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bytesizeCmd.Flags().StringVarP(&byteSizeF, "file", "f", "", "文件路径")
}
