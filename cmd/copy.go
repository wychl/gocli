/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package cmd

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var (
	copyFile string
	copyURL  string
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "复制",
	Long:  `复制`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := getCopyData(cmd, args)
		if err != nil {
			log.Fatal(err)
		}
		data = bytes.TrimSpace(data)
		if err := clipboard.WriteAll(string(data)); err != nil {
			log.Fatalln(err)
		}
	},
}

func getCopyData(cmd *cobra.Command, args []string) ([]byte, error) {
	if file := cmd.Flags().Lookup("file").Value.String(); file != "" {
		return os.ReadFile(file)
	}
	if u := cmd.Flags().Lookup("url").Value.String(); u != "" {
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	}
	if len(args) > 0 {
		return []byte(args[0]), nil
	}
	return io.ReadAll(os.Stdin)
}

func init() {
	rootCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	copyCmd.Flags().StringVarP(&copyFile, "file", "f", "", "文件路径")
	copyCmd.Flags().StringVarP(&copyURL, "url", "u", "", "文件网页链接")
}
