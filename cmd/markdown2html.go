/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application gocli.
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	formatterhtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var (
	mkURL      string
	mkFile     string
	htmlOutput string
)

const markdown2htmlExp = `
# markdown文件生成html
gocli markdown2html -f README.md
gocli markdown2html -f ${HOME}/README.md

# 根据markdown网页链接生成html
gocli markdown2html -u https://raw.githubusercontent.com/wychl/gocli/main/README.md

# 指定输出的文件名
gocli markdown2html -u https://raw.githubusercontent.com/wychl/gocli/main/README.md -o index.html

# unix管道方式生成html
cat README.md | gocli markdown2html
`

// markdown2htmlCmd represents the markdown2html command
var markdown2htmlCmd = &cobra.Command{
	Use:     "markdown2html",
	Short:   "markdown生成html",
	Long:    `markdown生成html`,
	Example: markdown2htmlExp,
	Run: func(cmd *cobra.Command, args []string) {
		var mkContent []byte
		switch {
		case mkURL != "":
			resp, err := http.Get(mkURL)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			if 200 <= resp.StatusCode && resp.StatusCode < 400 {
				mkContent, err = io.ReadAll(resp.Body) //nolint
				if err != nil {
					log.Fatal(err)
				}
			}
		case mkFile != "":
			var err error
			mkContent, err = os.ReadFile(mkFile) //nolint
			if err != nil {
				log.Fatal(err)
			}
		case mkFile == "" && mkURL == "":
			var err error
			mkContent, err = io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
		}
		md := goldmark.New(
			goldmark.WithExtensions(extension.GFM,
				extension.Table,
				extension.Strikethrough,
				extension.Linkify,
				extension.TaskList,
				extension.Typographer,
				extension.DefinitionList,
				extension.Footnote,
				highlighting.NewHighlighting(
					highlighting.WithStyle("monokai"),
					highlighting.WithFormatOptions(
						formatterhtml.Standalone(true),
						formatterhtml.WithAllClasses(true),
						formatterhtml.WithLineNumbers(true),
					),
				)),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
				parser.WithBlockParsers(),
				parser.WithInlineParsers(),
				parser.WithAttribute(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
				html.WithUnsafe(),
			),
		)
		var buf bytes.Buffer
		if err := md.Convert(mkContent, &buf); err != nil {
			panic(err)
		}
		if _, err := os.Stdout.Write(buf.Bytes()); err != nil {
			log.Fatal(err)
		}
		if htmlOutput != "" {
			if err := os.WriteFile(htmlOutput, buf.Bytes(), os.ModePerm); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", htmlOutput)
		}
	},
}

func init() {
	rootCmd.AddCommand(markdown2htmlCmd)
	markdown2htmlCmd.Flags().StringVarP(&mkURL, "url", "u", "", "markdown文件网页链接")
	markdown2htmlCmd.Flags().StringVarP(&mkFile, "file", "f", "", "markdown文件路径")
	markdown2htmlCmd.Flags().StringVarP(&htmlOutput, "output", "o", "", "输出的html文件路径")
}
