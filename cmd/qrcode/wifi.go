/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/spf13/cobra"
)

var (
	wifiEncrypType string // 加密类型 None WAP/WAP2 WEP
	wifiName       string // wifi名称
	wifiPass       string // wifi密码
	wifiIsHidden   bool   // 是否隐藏式wifi
	example        = "WIFI:T:WPA;S:test;P:root;H:true;;"
	wifiExp        = `
	# 生成wifi二维码
	gocli qrcode wifi --type=WAP --name=admin --pass=admin

	# 指定输出文件名
	gocli qrcode wifi --type=WAP --name=admin --pass=admin --output=qrcode.png

	# 指定尺寸
	gocli qrcode  wifi --type=WAP --name=admin --pass=admin --size=256

	参数：
	- type：加密类型，支持三种类型WAP None WEP
	- hidden：是否隐藏式，默认false
	- name：WIFI名称
	- pass：WIFI密码
	`
)

// wifiCmd represents the wifi command
var wifiCmd = &cobra.Command{
	Use:     "wifi",
	Short:   "生成wifi二维码",
	Long:    `生成wifi二维码`,
	Example: wifiExp,
	Run: func(cmd *cobra.Command, args []string) {
		hidden := ""
		if wifiIsHidden {
			hidden = "true"
		}
		content := fmt.Sprintf("WIFI:T:%s;S:%s;P:%s;H:%s;;",
			getWifiType(),
			wifiName,
			wifiPass,
			hidden)

		w := qrcode.NewQRCodeWriter()
		hints := make(map[gozxing.EncodeHintType]interface{})
		img, err := w.Encode(content, gozxing.BarcodeFormat_QR_CODE, qrcodeSize, qrcodeSize, hints)
		if err != nil {
			log.Fatal(err)
		}
		buf := new(bytes.Buffer)
		ext := "png"
		if list := strings.Split(qrcodeFile, "."); len(list) > 1 {
			ext = list[1]
		}
		switch ext {
		case "png":
			if err := png.Encode(buf, img); err != nil {
				log.Fatal(err)
			}
		case "jpeg", "jpg":
			if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: jpegQuality}); err != nil {
				log.Fatal(err)
			}
		}
		// 写入文件
		if err := os.WriteFile(qrcodeOutput, buf.Bytes(), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	qrcodeCmd.AddCommand(wifiCmd)

	// wifi账号参数
	wifiCmd.Flags().StringVarP(&wifiEncrypType, "type", "t", "none", "加密类型")
	wifiCmd.Flags().StringVarP(&wifiName, "name", "n", "", "账号")
	wifiCmd.Flags().StringVarP(&wifiPass, "pass", "p", "", "密码")
	wifiCmd.Flags().BoolVarP(&wifiIsHidden, "hidden", "d", false, "是否隐藏")

	// 输出二维码参数
	wifiCmd.Flags().StringVarP(&qrcodeOutput, "output", "o", "wifi.png", "输出的文件路径")
	wifiCmd.Flags().IntVarP(&qrcodeSize, "size", "s", 256, "二维码尺寸")
	wifiCmd.Flags().IntVarP(&jpegQuality, "quality", "q", 100, "jpeg图片质量参数")
}

func getWifiType() string {
	wifiType := strings.ToLower(wifiEncrypType)
	if strings.Contains(wifiType, "wap") {
		wifiType = "wap"
	}
	return strings.ToTitle(wifiType)
}
