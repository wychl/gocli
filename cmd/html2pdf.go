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

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

var (
	htmlURL   string
	htmlFile  string
	pdfOutput string
)

const html2pdfExp = `
# html文件转pdf
gocli html2pdf -f index.html
gocli html2pdf -f ${HOME}/index.html

# 网页链接转pdf
gocli html2pdf -u https://github.com/trending

# 指定输出的pdf文件名
gocli html2pdf -u https://github.com/trending -o github.pdf
`

// html2pdfCmd represents the html2pdf command
var html2pdfCmd = &cobra.Command{
	Use:   "html2pdf",
	Short: "html转pdf",
	Long:  `html转pdf`,
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
		if err := chromedp.Run(ctx, printToPDF(fp, &buf)); err != nil {
			log.Fatal(err)
		}
		if err := os.MkdirAll(filepath.Dir(pdfOutput), os.ModePerm); err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile(pdfOutput, buf, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", pdfOutput)
	},
}

func init() {
	rootCmd.AddCommand(html2pdfCmd)
	html2pdfCmd.Flags().StringVarP(&htmlURL, "url", "u", "", "网页链接")
	html2pdfCmd.Flags().StringVarP(&htmlFile, "file", "f", "", "html文件路径")
	html2pdfCmd.Flags().StringVarP(&pdfOutput, "output", "o", "output.pdf", "输出的文件")
}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
