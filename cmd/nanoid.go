/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/jaevor/go-nanoid"
	"github.com/spf13/cobra"
)

var (
	nanoidSize     int
	nanoidAlphabet string
	nanoidExp      = `
	# 生成指定长度的id
	gocli nanoid --s 15

	# 生成指定长度和字符集的ID
	gocli nanoid --size 10 --alphabet abc
	`
)

// nanoidCmd represents the nanoid command
var nanoidCmd = &cobra.Command{
	Use:     "nanoid",
	Short:   "生成随机ID",
	Long:    `生成随机ID`,
	Example: nanoidExp,
	Run: func(cmd *cobra.Command, args []string) {
		var generator func() string
		var err error
		if nanoidAlphabet != "" {
			generator, err = nanoid.CustomASCII(nanoidAlphabet, nanoidSize) //nolint
			if err != nil {
				log.Fatal(err)
			}
		} else {
			generator, err = nanoid.Standard(nanoidSize) //nolint
		}
		os.Stdout.WriteString(generator())
	},
}

func init() {
	rootCmd.AddCommand(nanoidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nanoidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nanoidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	nanoidCmd.Flags().IntVarP(&nanoidSize, "size", "s", 15, "Generated ID size")
	nanoidCmd.Flags().StringVarP(&nanoidAlphabet, "alphabet", "a", "", "Alphabet to use")
}
