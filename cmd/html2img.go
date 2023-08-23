/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

const img2pdfExp = `
# img文件转图片
gocli html2img -f index.html
gocli html2img -f ${HOME}/index.html

# 网页链接转图片
gocli html2img -u https://github.com/trending
`

// html2ImgCmd represents the html2img command
var html2ImgCmd = &cobra.Command{
	Use:   "html2img",
	Short: "html装图片",
	Long:  `html装图片`,
	Run: func(cmd *cobra.Command, args []string) {
		fp := htmlURL
		if fp == "" {
			fp = htmlFile
			if fp != "" {
				absF, err := filepath.Abs(fp)
				if err != nil {
					log.Fatal(err)
				}
				fp = fmt.Sprintf("file://%s", absF) //nolint
			}
			if fp == "" {
				data, err := io.ReadAll(os.Stdin)
				if err != nil {
					log.Fatal(err)
				}
				tmpf := fmt.Sprintf("%d-index.html", time.Now().Unix())
				absF, err := filepath.Abs(tmpf)
				if err != nil {
					log.Fatal(err)
				}
				fp = fmt.Sprintf("file://%s", absF)
				if err := os.WriteFile(absF, data, os.ModePerm); err != nil {
					log.Fatal(err)
				}
				defer os.Remove(tmpf)
			}
		}
		if fp == "" {
			return
		}
		// create context
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()
		// capture pdf
		var buf []byte
		if err := chromedp.Run(ctx, fullScreenshot(fp, 90, &buf)); err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile("screenshot.png", buf, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", "screenshot.png")
	},
}

func init() {
	rootCmd.AddCommand(html2ImgCmd)

	html2ImgCmd.Flags().StringVarP(&htmlURL, "url", "u", "", "网页链接")
	html2ImgCmd.Flags().StringVarP(&htmlFile, "file", "f", "", "img文件路径")
}

// fullScreenshot takes a screenshot of the entire browser viewport.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
