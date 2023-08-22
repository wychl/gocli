/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

var jsonEscape = false

// htmlminifyCmd represents the htmlminify command
var htmlminifyCmd = &cobra.Command{
	Use:   "htmlminify",
	Short: "html minify",
	Long:  `html minify`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var data []byte
		if htmlFile != "" {
			data, err = os.ReadFile("htmlFile")
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		if htmlURL != "" {
			resp, err := http.Get(htmlURL)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer resp.Body.Close()
			if 200 <= resp.StatusCode && resp.StatusCode < 400 {
				data, err = io.ReadAll(resp.Body) //nolint
				if err != nil {
					log.Fatal(err)
					return
				}
			}
		}
		if len(data) == 0 && len(args) > 0 {
			data = []byte(args[0])
		}
		if len(data) == 0 {
			data, err = io.ReadAll(os.Stdin) //nolint
			if err != nil {
				log.Fatal(err)
				return
			}
		}

		if jsonEscape {
			buf := new(bytes.Buffer)
			e := json.NewEncoder(buf)
			e.SetEscapeHTML(false)
			if err := e.Encode(string(data)); err != nil {
				log.Fatal(err)
				return
			}
			data = buf.Bytes()
		}
		m := minify.New()
		m.AddFunc("text/html", html.Minify)
		b, err := m.Bytes("text/html", data)
		if err != nil {
			log.Fatal(err)
			return
		}
		b = bytes.TrimPrefix(b, []byte(`"`))
		b = bytes.TrimSuffix(b, []byte(`"`))
		os.Stdout.Write(b)
	},
}

func init() {
	rootCmd.AddCommand(htmlminifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// htmlminifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// htmlminifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	htmlminifyCmd.Flags().BoolVarP(&jsonEscape, "jsonescape", "e", false, "将html字符串转义为json字符串")
}
