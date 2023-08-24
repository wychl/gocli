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

	"github.com/ipipdotnet/ipdb-go"
	"github.com/spf13/cobra"
	"github.com/wychl/gocli/asset"
)

const geoIPExp = `
# 解析IP
gocli geoip 124.78.251.238

# 管道方式解析IP
echo 124.78.251.238 gocli geoip
`

// geoipCmd represents the geoip command
var geoipCmd = &cobra.Command{
	Use:     "geoip",
	Short:   "IP地址",
	Long:    `IP地址`,
	Example: geoIPExp,
	Run: func(cmd *cobra.Command, args []string) {
		var ipByte []byte
		if len(args) > 0 {
			ipByte = []byte(args[0])
		}
		if len(ipByte) == 0 {
			var err error
			ipByte, err = io.ReadAll(os.Stdin) //nolint
			if err != nil {
				log.Fatal(err)
			}
		}
		ipStr := string(ipByte)
		ipStr = strings.TrimSpace(ipStr)
		db, err := ipdb.NewCityFromBytes(asset.QQWRY)
		if err != nil {
			log.Fatal(err)
		}
		info, err := db.FindInfo(ipStr, "CN")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("国家：%s\n城市：%s\n",
			info.CountryName,
			info.CityName)
	},
}

func init() {
	rootCmd.AddCommand(geoipCmd)
}
