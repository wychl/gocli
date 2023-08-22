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
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var tsZone = ""

const ts2DateExp = `
# 时间戳转日期
gocli ts2date 1694102400

# 13位的时间戳转日期
gocli ts2date 1694102400000

# 指定时区
gocli ts2date 1694102400000 --zone=UTC

# 以管道方式执行命令
echo "1694102400" | gocli ts2date
`

// ts2DateCmd represents the tstodate command
var ts2DateCmd = &cobra.Command{
	Use:     "ts2date",
	Short:   "时间戳转日期",
	Long:    `时间戳转日期`,
	Example: ts2DateExp,
	Run: func(cmd *cobra.Command, args []string) {
		var ts []byte
		if len(args) > 0 {
			ts = []byte(args[0])
		}
		if len(ts) == 0 {
			var err error
			ts, err = io.ReadAll(os.Stdin) //nolint
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

		tsStr := string(ts)
		tsStr = strings.TrimSuffix(tsStr, "\n")
		tsNum, err := strconv.ParseInt(tsStr, 0, 0)
		if err != nil {
			log.Fatal(err)
		}
		if len(tsStr) == 10 {
			fmt.Println(time.Unix(tsNum, 0).In(zone))
		} else {
			fmt.Println(time.UnixMilli(tsNum).In(zone))
		}
	},
}

func init() {
	rootCmd.AddCommand(ts2DateCmd)
	ts2DateCmd.Flags().StringVarP(&tsZone, "zone", "z", "", "时区,默认本地时区")
}
