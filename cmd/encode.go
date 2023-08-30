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
	"io"
	"log"
	"os"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/spf13/cobra"
)

var (
	qrcodeOutput string
	qrcodeSize   int
	jpegQuality  = 100
	encodeExp    = `
# 生成二维码
gocli qrcode encode "hello wrod"

# 指定输出文件名
gocli qrcode encode "hello wrod" --output=qrcode.png

# 指定尺寸
gocli qrcode encode "hello wrod" --size=256
`
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:     "encode",
	Short:   "生成二维码",
	Long:    `生成二维码`,
	Example: encodeExp,
	Run: func(cmd *cobra.Command, args []string) {
		var dataBytes []byte
		if len(args) > 0 {
			dataBytes = []byte(args[0])
		}
		file := os.Stdin
		fi, err := file.Stat()
		if err != nil {
			fmt.Println("file.Stat()", err)
		}
		if len(dataBytes) == 0 && fi.Size() > 0 {
			var err error
			dataBytes, err = io.ReadAll(os.Stdin) //nolint
			if err != nil {
				log.Fatal(err)
			}
		}
		dataBytes = bytes.TrimSpace(dataBytes)
		w := qrcode.NewQRCodeWriter()
		hints := make(map[gozxing.EncodeHintType]interface{})
		img, err := w.Encode(string(dataBytes), gozxing.BarcodeFormat_QR_CODE, qrcodeSize, qrcodeSize, hints)
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
	qrcodeCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().StringVarP(&qrcodeOutput, "output", "o", "qrcode.png", "输出的文件路径")
	encodeCmd.Flags().IntVarP(&qrcodeSize, "size", "s", 256, "二维码尺寸")
	encodeCmd.Flags().IntVarP(&jpegQuality, "quality", "q", 100, "jpeg图片质量参数")
}
