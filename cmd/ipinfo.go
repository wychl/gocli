/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

const ipinfoExp = `
gocli ipinfo
`

// ipinfoCmd represents the ipinfo command
var ipinfoCmd = &cobra.Command{
	Use:     "ipinfo",
	Short:   "外网的IP地址信息",
	Long:    `本地外网IP地址信息`,
	Example: ipinfoExp,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://ipinfo.io", nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		var result struct {
			IP       string `json:"ip"`
			Hostname string `json:"hostname"`
			City     string `json:"city"`
			Region   string `json:"region"`
			Country  string `json:"country"`
			Loc      string `json:"loc"`
			Org      string `json:"org"`
			Timezone string `json:"timezone"`
			Readme   string `json:"readme"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&result)
		fmt.Printf("IP：%s\n国家:%s\n城市：%s\n组织：%s\n经纬度：%s\n时区：%s\n",
			result.IP,
			result.Country,
			result.City,
			result.Org,
			result.IP,
			result.Timezone)
	},
}

func init() {
	rootCmd.AddCommand(ipinfoCmd)
}
