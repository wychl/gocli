/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wychl/gocli/command/cron"
)

var (
	cronWithSecond bool        // 是否包含秒
	cronZone       string = "" // 时区
	cronStart      string = "" // 开始日期
)

const cronExp = `
# 解析5位cron表达式
gocli cron '0 */2 * * *'

# 解析6位cron表达式
gocli cron '*/5 * * * *'

# 指定时区
gocli cron '*/5 * * * *' --zone=UTC
`

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:     "cron",
	Short:   "解析cron表达式",
	Long:    `解析cron表达式`,
	Example: cronExp,
	Run: func(cmd *cobra.Command, args []string) {
		var spec []byte
		if len(args) > 0 {
			spec = []byte(args[0])
		}
		if len(spec) == 0 {
			var err error
			spec, err = io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
		}
		parser := cron.New()
		parser.Run(&cron.Config{Exp: string(spec), Zone: cronZone, Start: cronStart, WithSecond: cronWithSecond})
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
	cronCmd.Flags().BoolVarP(&cronWithSecond, "second", "s", false, "是否包含秒")
	cronCmd.Flags().StringVarP(&cronZone, "zone", "z", "", "时区,默认本地时区")
	cronCmd.Flags().StringVarP(&cronStart, "start", "", "", "开始日期，格式：2006-01-02")
}
