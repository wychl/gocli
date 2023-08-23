/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
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
		var parser cron.Parser
		if cronWithSecond {
			parser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
		} else {
			parser = cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
		}
		schedule, err := parser.Parse(string(spec))
		if err != nil {
			log.Fatal(err)
		}
		zone := time.Local
		if cronZone != "" {
			var err error
			zone, err = time.LoadLocation(cronZone) //nolint
			if err != nil {
				log.Fatal(err)
			}
		}
		startDate := time.Now()
		if cronStart != "" {
			startDate, err = time.Parse(time.DateOnly, cronStart)
			if err != nil {
				log.Fatal(err)
			}
		}
		builder := strings.Builder{}
		cur := startDate.In(zone)
		for i := 0; i < 10; i++ {
			cur = schedule.Next(cur).In(zone)
			builder.WriteString(fmt.Sprintf("%s\n", cur))
		}
		fmt.Println(builder.String())
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
	cronCmd.Flags().BoolVarP(&cronWithSecond, "second", "s", false, "是否包含秒")
	cronCmd.Flags().StringVarP(&cronZone, "zone", "z", "", "时区,默认本地时区")
	cronCmd.Flags().StringVarP(&cronStart, "start", "", "", "开始日期，格式：2006-01-02")
}
