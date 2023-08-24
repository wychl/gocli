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
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

const curl2CodeExp = `
注意： 此命令需要联网
网址：https://curlconverter.com/

# curl命令生成代码
gocli curl2code

# 指定生成的语言
gocli curl2code -l=rust


# 指定curl命令代码
gocli curl2code -l=rust -c="curl 'http://fiddle.jshell.net/echo/html/' \
    -H 'Origin: http://fiddle.jshell.net' \
    -H 'Accept-Encoding: gzip, deflate' \
    -H 'Accept-Language: en-US,en;q=0.8' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36' \
    -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' \
    -H 'Accept: */*' \
    -H 'Referer: http://fiddle.jshell.net/_display/' \
    -H 'X-Requested-With: XMLHttpRequest' \
    -H 'Connection: keep-alive' \
    --data 'msg1=wow&msg2=such&msg3=data'"

# 管道指定curl命令代码
echo "curl 'http://fiddle.jshell.net/echo/html/' \
    -H 'Origin: http://fiddle.jshell.net' \
    -H 'Accept-Encoding: gzip, deflate' \
    -H 'Accept-Language: en-US,en;q=0.8' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36' \
    -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' \
    -H 'Accept: */*' \
    -H 'Referer: http://fiddle.jshell.net/_display/' \
    -H 'X-Requested-With: XMLHttpRequest' \
    -H 'Connection: keep-alive' \
    --data 'msg1=wow&msg2=such&msg3=data' --compressed" | gocli curl2code -l=rust

#支持的语言列表：
- ansible
- cfml
- clojure
- csharp
- dart
- elixir
- go
- har
- http
- httpie
- java, java-httpurlconnection, java-jsoup, java-okhttp
- javascript, javascript-jquery, javascript-xhr
- json
- kotlin
- matlab
- node, node-http, node-axios, node-got, node-ky, node-request, node-superagent
- ocaml
- php, php-guzzle, php-requests
- powershell, powershell-webrequest
- python (the default)
- r
- ruby
- rust
- swift
- wget
`

var (
	curlCode = "curl http://example.com" // curl命令
	curlLang = "go"                      // 生成的语言
)

// curl2codeCmd represents the curl2code command
var curl2codeCmd = &cobra.Command{
	Use:     "curl2code",
	Short:   "curl转代码",
	Long:    `curl转代码`,
	Example: curl2CodeExp,
	Run: func(cmd *cobra.Command, args []string) {
		var curlCodeBytes []byte
		if len(args) > 0 {
			curlCodeBytes = []byte(args[0])
		}
		file := os.Stdin
		fi, err := file.Stat()
		if err != nil {
			fmt.Println("file.Stat()", err)
		}
		if len(curlCodeBytes) == 0 && fi.Size() > 0 {
			var err error
			curlCodeBytes, err = io.ReadAll(os.Stdin) //nolint
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(curlCodeBytes) > 0 {
			curlCode = string(curlCodeBytes)
		}
		curlCode = strings.TrimSpace(curlCode)
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()
		fp := fmt.Sprintf("https://curlconverter.com/%s", curlLang)
		var code string
		if err := chromedp.Run(ctx, toCode(fp, curlCode, &code)); err != nil {
			log.Fatal(err)
		}
		os.Stdout.WriteString(code)
	},
}

func init() {
	rootCmd.AddCommand(curl2codeCmd)
	curl2codeCmd.Flags().StringVarP(&curlCode, "curl", "c", "curl http://example.com", "curl命令")
	curl2codeCmd.Flags().StringVarP(&curlLang, "language", "l", "go", "生成的语言")
}

func toCode(urlstr, curlCode string, text *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.SetValue("#curl-code", curlCode, chromedp.ByID),
		chromedp.Text(`#generated-code`, text, chromedp.ByID, chromedp.NodeVisible),
	}
}
