/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var tsSize int

const date2TsExp = `
# 时间转时间戳
gocli date2ts '2023-09-08 00:00:00'

# 生成13位的时间戳
gocli date2ts '2023-09-08 00:00:00' --size=13

# 指定时区
gocli date2ts '2023-09-08 00:00:00' --zone=UTC

# 以管道方式执行命令
echo "2023-09-08 00:00:00" | gocli date2ts
`

// date2TsCmd represents the datetots command
var date2TsCmd = &cobra.Command{
	Use:     "date2ts",
	Short:   "日期转时间戳",
	Long:    `日期转时间戳`,
	Example: date2TsExp,
	Run: func(cmd *cobra.Command, args []string) {
		var dateBytes []byte
		if len(args) > 0 {
			dateBytes = []byte(args[0])
		}
		if len(dateBytes) == 0 {
			var err error
			dateBytes, err = io.ReadAll(os.Stdin) //nolint
			if err != nil {
				log.Fatal(err)
			}
		}
		zone := time.Local
		if tsZone != "" {
			var err error
			zone, err = time.LoadLocation(tsZone) //nolint
			if err != nil {
				log.Fatal(err)
			}
		}
		dateStr := string(dateBytes)
		dateStr = strings.TrimSuffix(dateStr, "\n")
		date, err := time.ParseInLocation(time.DateTime, dateStr, zone)
		if err != nil {
			log.Fatal(err)
		}
		if tsSize == 10 {
			fmt.Println(date.Unix())
		} else {
			fmt.Println(date.UnixMilli())
		}
	},
}

func init() {
	rootCmd.AddCommand(date2TsCmd)
	date2TsCmd.Flags().StringVarP(&tsZone, "zone", "z", "", "时区,默认本地时区")
	date2TsCmd.Flags().IntVarP(&tsSize, "size", "s", 10, "时间戳长度")
}
