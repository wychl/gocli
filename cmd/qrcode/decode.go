/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/spf13/cobra"
)

var (
	qrcodeFile string
	qrcodeURL  string
	qrcodeExp  = `
# 解析本地图片
gocli qrcode decode --file=qrcode.png

# 解析网络图片
gocli qrcode decode --url=http://example.com/qrcode.png 	
	`
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:     "decode",
	Short:   "解析二维码内容",
	Long:    `解析二维码内容`,
	Example: qrcodeExp,
	Run: func(cmd *cobra.Command, args []string) {
		var r io.Reader
		if qrcodeFile != "" {
			file, err := os.Open(qrcodeFile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			r = file
		}
		if qrcodeURL != "" {
			resp, err := http.Get(qrcodeURL)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			r = resp.Body
		}

		img, _, _ := image.Decode(r)
		bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
		qrReader := qrcode.NewQRCodeReader()
		result, _ := qrReader.DecodeWithoutHints(bmp)
		fmt.Print(result)
	},
}

func init() {
	qrcodeCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().StringVarP(&qrcodeFile, "file", "f", "", "二维码文件路径")
	decodeCmd.Flags().StringVarP(&qrcodeURL, "url", "u", "", "二维码文件URL地址")
}
